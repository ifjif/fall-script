package module

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"zzc/fall-script/src/ast"
	"zzc/fall-script/src/ast/marshal"
	binarychunck "zzc/fall-script/src/binary_chunck"
	"zzc/fall-script/src/builtin/vmb"
	"zzc/fall-script/src/compiler"
	"zzc/fall-script/src/ir"
	"zzc/fall-script/src/macro"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/parser"
	"zzc/fall-script/src/semantic"
	"zzc/fall-script/src/utils"
)

type (
	FileType byte
)

const (
	SOURCE FileType = iota
	BYTECODE
)

const (
	BuildDir       = ".build"
	SourceSuffix   = ".fs"
	BytecodeSuffix = ".fsc"
	MetaSuffix     = ".fsm"
)

type Loader struct {
	globalSymbol   *ir.SymbolTable
	sourceSuffix   string
	bytecodeSuffix string
	metaSuffix     string
	buildDir       string
	Build          bool
}

func NewLoader() *Loader {
	globalSt := ir.NewSymbolTable()
	for i, fn := range vmb.Builtins {
		globalSt.DefineBuiltin(i, fn.Name)
	}

	loader := &Loader{
		globalSymbol:   globalSt,
		sourceSuffix:   SourceSuffix,
		bytecodeSuffix: BytecodeSuffix,
		metaSuffix:     MetaSuffix,
		buildDir:       BuildDir,
	}
	return loader
}

func (l *Loader) CreateDir(file string) {
	dir := filepath.Dir(file)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		panic(err)
	}
}

func (l *Loader) BuildFilepath(file string) string {
	return filepath.Join(l.buildDir, file)
}

func (l *Loader) BytecodeFilepath(file string) string {
	return file + l.bytecodeSuffix
}

func (l *Loader) SourceFilepath(file string) string {
	return file + l.sourceSuffix
}

func (l *Loader) MetaFilepath(file string) string {
	return file + l.metaSuffix
}

func (l *Loader) readFile(file string) ([]byte, FileType) {
	data, err := l.readBytecodeFile(file)
	if err == nil {
		return data, BYTECODE
	}
	data, err = l.readSourceFile(file)
	if err != nil {
		// todo没有发现源或字节码文件
		panic(err)
	}
	return data, SOURCE
}

func (l *Loader) readBytecodeFile(file string) ([]byte, error) {
	nfile := l.BytecodeFilepath(file)
	data, err := os.ReadFile(nfile)
	return data, err
}

func (l *Loader) readSourceFile(file string) ([]byte, error) {
	nfile := l.SourceFilepath(file)
	data, err := os.ReadFile(nfile)

	return data, err
}

func (l *Loader) readMetaFile(file string) ([]byte, error) {
	nfile := l.MetaFilepath(file)
	data, err := os.ReadFile(nfile)

	return data, err
}

func (l *Loader) GenerateAST(file string) *ast.Program {
	data, err := l.readSourceFile(file)
	if err != nil {
		panic(err)
	}

	return l.parse(data)
}

func (l *Loader) parse(data []byte) *ast.Program {
	p := parser.NewParser(string(data))
	program := p.Parse()
	if len(p.Errors()) != 0 {
		error := strings.Join(p.Errors(), "\n")
		panic(error)
	}

	return program
}

func (l *Loader) LoadMetaFile(file string) (ExportMetas, error) {
	// 如是 build 从 .build中加载
	if l.Build {
		file = l.BuildFilepath(file)
	}
	data, err := l.readMetaFile(file)
	if err != nil {
		return nil, err
	}

	fr := binarychunck.FlReader{}
	fr.InitData(data)
	exportMetas := ExportMetas{}
	count := fr.ReadUint16()

	// Methods    []string // struct AST
	// FieldTotal int      // struct AST
	// Fields     []int    // struct AST 所有字段的offset
	for range count {
		name := fr.ReadStr()
		source := fr.ReadStr()
		imported := fr.ReadStr()
		exportIdx := fr.ReadUint16()
		origin := fr.ReadUint8()

		methodCount := fr.ReadVarint()
		methods := make([]string, methodCount)
		for i := range methodCount {
			methods[i] = fr.ReadStr()
		}

		fieldTotal := fr.ReadVarint()

		fieldCount := fr.ReadVarint()
		fields := make([]int, fieldCount)
		for i := range fieldCount {
			fields[i] = int(fr.ReadVarint())
		}

		ast, ndata := marshal.Unmarshal(fr.Bytes())
		fr.InitData(ndata)

		em := &ExportMeta{
			Name:       name,
			Source:     source,
			ExportIdx:  int(exportIdx),
			Imported:   imported,
			Ast:        ast,
			Origin:     ExportOrigin(origin),
			Methods:    methods,
			FieldTotal: int(fieldTotal),
			Fields:     fields,
		}

		exportMetas[name] = em
	}

	return exportMetas, nil
}

// 先 字节码 再 源码
func (l *Loader) LoadFile(file string) *object.Module {
	data, ft := l.readFile(file)
	if ft == BYTECODE {
		return l.LoadBytecode(data)
	}
	return l.LoadText(data, file)
}

func (l *Loader) LoadText(data []byte, file string) *object.Module {
	program := l.parse(data)

	env := object.NewEnvironment()
	// imports, exports
	program, imports, exports, importMetas := ResolveMacrosFromProgram(l, program, file, env)
	//	fmt.Println(env.Inspect())
	//	fmt.Println("代码块：")
	//	fmt.Println(program.String())
	//	fmt.Println("去宏后的import: ")
	//	for _, imp := range imports {
	//		fmt.Println(imp.String())
	//	}
	//	fmt.Println("去宏后的exports: ")
	//	for _, exp := range exports {
	//		fmt.Println(exp.Declaration.String())
	//	}
	nProgram := macro.ExpandMacros(program, env)
	np := nProgram.(*ast.Program)
	// fmt.Println("展开后的代码块：")
	// fmt.Println(nProgram.String())
	np.Imports = imports
	np.Exports = exports

	// 语义分析
	analyzer := semantic.NewAnalyzer(l.globalSymbol)
	ip := analyzer.Analyze(np, importMetas)
	if len(analyzer.Errors()) > 0 {
		fmt.Println(strings.Join(analyzer.Errors(), "\n"))
		fmt.Println()
	}
	// fmt.Println(ip.String())

	// cmp := compiler.NewCompiler(nProgram, imports, exports, l.rootSt)
	// cmp.SetStructAsts(program.Structs)
	// cmp.SetMethods(program.Methods)
	// cmp.SetPromotedFns(program.PromotedFns)
	// cmp.Compile()
	cmp := compiler.NewCompiler(ip)
	cmp.Compiler()
	module := cmp.MainModule()
	module.Name = file
	utils.PrintModule(module, "")
	return module
}

func (l *Loader) BuildProgram(prev, file string) {
	source := ResolveImportPath(prev, file)

	data, err := l.readSourceFile(source)
	if err != nil {
		panic(err)
	}

	l.doBuild(data, source)
}

func (l *Loader) doBuild(data []byte, file string) {
	program := l.parse(data)

	env := object.NewEnvironment()
	program, imports, exports, importMetas := ResolveMacrosFromProgram(l, program, file, env)
	nProgram := macro.ExpandMacros(program, env)

	np := nProgram.(*ast.Program)
	np.Imports = imports
	np.Exports = exports

	// 语义分析
	analyzer := semantic.NewAnalyzer(l.globalSymbol)
	ip := analyzer.Analyze(np, importMetas)
	if len(analyzer.Errors()) > 0 {
		err := strings.Join(analyzer.Errors(), "\n")
		panic(err)
	}

	cmp := compiler.NewCompiler(ip)
	cmp.Compiler()

	module := cmp.MainModule()
	module.Name = file

	bc := binarychunck.Dump(module)
	bf := l.BuildFilepath(l.BytecodeFilepath(file))

	l.CreateDir(bf)
	os.WriteFile(bf, bc, 0o644)

	for _, imp := range imports {
		l.BuildProgram(file, imp.Source)
	}
}

func (l *Loader) LoadBytecode(data []byte) *object.Module {
	mo := binarychunck.Undump(data)
	return mo
}

func (l *Loader) DumpFile(file string) {
	input, err := l.readSourceFile(file)
	if err != nil {
		panic(err)
	}

	mo := l.LoadText(input, file)

	data := binarychunck.Dump(mo)
	// fmt.Println("DUMP-------------------------------------------")
	// fmt.Printf("%v\n", data)
	outName := l.BytecodeFilepath(file)
	os.WriteFile(outName, data, 0o644)
	// fmt.Println("UNDUMP-------------------------------------------")
	l.LoadBytecode(data)
	// utils.PrintModule(nmo, "")
}

func (l *Loader) DumpMeta(file string, data []byte) {
	mf := l.MetaFilepath(file)
	if l.Build {
		mf = l.BuildFilepath(mf)
	}
	l.CreateDir(mf)
	os.WriteFile(mf, data, 0o644)
}
