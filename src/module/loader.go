package module

import (
	"encoding/binary"
	"fmt"
	"os"
	"strings"

	"zzc/fall-script/src/ast"
	"zzc/fall-script/src/ast/marshal"
	binarychunck "zzc/fall-script/src/binary_chunck"
	"zzc/fall-script/src/builtin/vmb"
	"zzc/fall-script/src/compiler"
	"zzc/fall-script/src/macro"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/parser"
)

type FileType byte

const (
	SOURCE FileType = iota
	BYTECODE
)

const (
	SourceSuffix   = ".fs"
	BytecodeSuffix = ".fsc"
	MetaSuffix     = ".fsm"
)

type Loader struct {
	rootSt         *compiler.SymbolTable
	sourceSuffix   string
	bytecodeSuffix string
	metaSuffix     string
}

func NewLoader() *Loader {
	st := compiler.NewSymbolTable()
	for i, fn := range vmb.Builtins {
		st.DefineBuiltin(i, fn.Name)
	}

	loader := &Loader{
		rootSt:         st,
		sourceSuffix:   SourceSuffix,
		bytecodeSuffix: BytecodeSuffix,
		metaSuffix:     MetaSuffix,
	}
	return loader
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
	data, err := l.readMetaFile(file)
	if err != nil {
		return nil, err
	}

	exportMetas := ExportMetas{}
	count := binary.BigEndian.Uint16(data)
	data = data[2:]

	for range count {
		nl := binary.BigEndian.Uint16(data)
		data = data[2:]
		name := string(data[:nl])
		data = data[nl:]
		sl := binary.BigEndian.Uint16(data)
		data = data[2:]
		source := string(data[:sl])
		data = data[sl:]
		il := binary.BigEndian.Uint16(data)
		data = data[2:]
		imported := string(data[:il])
		data = data[il:]
		exportIdx := binary.BigEndian.Uint16(data)
		data = data[2:]
		origin := data[0]
		data = data[1:]
		ast, nd := marshal.Unmarshal(data)
		data = nd

		em := &ExportMeta{
			Name:      name,
			Source:    source,
			ExportIdx: int(exportIdx),
			Imported:  imported,
			Ast:       ast,
			Origin:    ExportOrigin(origin),
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
	program, imports, exports := ResolveMacrosFromProgram(l, program, file, env)
	fmt.Println(env.Inspect())
	fmt.Println("代码块：")
	fmt.Println(program.String())
	fmt.Println("去宏后的import: ")
	for _, imp := range imports {
		fmt.Println(imp.String())
	}
	fmt.Println("去宏后的exports: ")
	for _, exp := range exports {
		fmt.Println(exp.Declaration.String())
	}
	nProgram := macro.ExpandMacros(program, env)
	fmt.Println("展开后的代码块：")
	fmt.Println(nProgram.String())

	cmp := compiler.NewCompiler(nProgram, imports, exports, l.rootSt)
	cmp.Compile()
	module := cmp.MainModule()
	module.Name = file
	//	utils.PrintModule(module, "")
	return module
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
