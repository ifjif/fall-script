package module

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"zzc/fall-script/src/ast"
	"zzc/fall-script/src/ast/marshal"
	"zzc/fall-script/src/macro"
)

/*
* 处理imports
* 根据 imports中信息，先去目标exports找
* 如果是标识符，先从ast中找，没有找到，再从imports中找，如果都没有，报错
*
*

	Token       token.Token
	Declaration Node

*
*/

type ExportOrigin byte

const (
	OriginLocal ExportOrigin = iota
	OriginReExport
)

// 自定义序列化 AST
type ExportMeta struct {
	Name       string
	Ast        ast.Node
	Source     string
	ExportIdx  int
	Origin     ExportOrigin
	Imported   string
	Methods    []string // struct ASt
	FieldTotal int      // struct AST
}

func NewExportMeta(name string, ast ast.Node, source string, exportIdx int) *ExportMeta {
	return &ExportMeta{
		Name:      name,
		Ast:       ast,
		ExportIdx: exportIdx,
		Source:    source,
		Origin:    OriginLocal,
	}
}

func (em *ExportMeta) SetReExportInfo(imported string) {
	em.Origin = OriginReExport
	em.Imported = imported
}

type ExportMetas map[string]*ExportMeta

func (ems ExportMetas) Serialize(file string) []byte {
	// fmt.Printf("====================序列化模块：%s\n======================", file)
	output := ExportMetas{}
	for name, em := range ems {
		nem := NewExportMeta(em.Name, nil, em.Source, em.ExportIdx)
		if em.Origin == OriginReExport {
			//	fmt.Printf("重导入：%s -> %s at %s", em.Name, em.Imported, em.Source)
			nem.SetReExportInfo(em.Imported)
			//	fmt.Printf("%+v\n", nem)
			output[name] = nem
			continue
		}

		fn, ok := em.Ast.(*ast.FnExpr)
		if !ok {
			nem.Ast = em.Ast
			output[name] = nem
			continue
		}

		ok, _ = macro.IsMacroDefinition(fn)
		nfn := &ast.FnExpr{
			Token:  fn.Token,
			Name:   fn.Name,
			Ident:  fn.Ident,
			Params: fn.Params,
			UnName: fn.UnName,
			Attrs:  fn.Attrs,
			Body:   nil,
		}
		if ok {
			nfn.Body = fn.Body
		}
		nem.Ast = nfn
		output[name] = nem
	}

	//	fmt.Println("=================================序列化ast")
	//	Name      string
	//	Ast       ast.Node
	//	Source    string
	//	ExportIdx int
	//	Origin    ExportOrigin
	//	Imported  string
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, uint16(len(output)))
	for _, em := range output {
		binary.Write(&buf, binary.BigEndian, uint16(len(em.Name)))
		buf.WriteString(em.Name)
		binary.Write(&buf, binary.BigEndian, uint16(len(em.Source)))
		buf.WriteString(em.Source)
		binary.Write(&buf, binary.BigEndian, uint16(len(em.Imported)))
		buf.WriteString(em.Imported)
		binary.Write(&buf, binary.BigEndian, uint16(em.ExportIdx))
		binary.Write(&buf, binary.BigEndian, uint8(em.Origin))
		data := marshal.MarshalAst(em.Ast)
		_, err := buf.Write(data)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}

	//	fmt.Println("=================================序列化ast end")
	return buf.Bytes()
}

type ExportMetasRegister map[string]ExportMetas

/*
* export 流程
* 取export中 非 identifier 的 加入ExportMetas
* 剩下的identifier
* 	1. 从AST中找,找到加入ExportMetas
* 	2. 还剩下，从import中找
* 		 看源是否在 ExportMetaRegister中：
* 		 	有，取出
* 		 	无，解析
*      寻找，如果没有，panic; 找到加入 ExportMetas
 */
func resolveExports(l *Loader, file string, er ExportMetasRegister) ExportMetas {
	exports2, ok := er[file]
	if ok {
		return exports2
	}

	// 从 meta 文件中加载
	exports2, err := l.LoadMetaFile(file)

	fmt.Println("===================================================加载meta")
	if err == nil {
		fmt.Println("===================================================来自meta")
		er[file] = exports2
		return exports2
	}
	// 从源文件中加载
	program := l.GenerateAST(file)

	imports, exports := CollectImportsAndExports(program)

	return resolveExports2(l, file, program, imports, exports, er)
}

func resolveExports2(l *Loader, file string, program *ast.Program, imports []*ast.ImportStmt, exports []*ast.ExportStmt, er ExportMetasRegister) ExportMetas {
	source := file

	structs := program.Structs
	// methods := program.Methods

	exports2 := ExportMetas{}
	ident := map[string]int{}
	exportStruct := map[string]*ast.StructDeclStmt{}

	for i, exp := range exports {
		node := exp.Declaration
		var name string
		var astNode ast.Node
		switch node := node.(type) {
		case *ast.LetStmt:
			name = node.Name.Value
			astNode = node.Value
		case *ast.FnExpr:
			if !node.UnName {
				name = node.Name
			}
			astNode = node
		case *ast.StructDeclStmt:
			exportStruct[node.Name.Value] = node
		case *ast.IdentExpr:
			ident[node.Value] = i
		}

		if astNode != nil {
			e := NewExportMeta(name, astNode, source, i)
			exports2[name] = e
		}
	}

	// 从 ast 中找
	for _, stmt := range program.Stmts {
		if len(ident) <= 0 {
			break
		}
		switch stmt := stmt.(type) {
		case *ast.LetStmt:
			name := stmt.Name.Value
			if idx, ok := ident[name]; ok {
				ast := stmt.Value
				e := NewExportMeta(name, ast, source, idx)
				exports2[name] = e
				delete(ident, name)
			}
		case *ast.ExprStmt:
			expr := stmt.Expr
			if expr, ok := expr.(*ast.FnExpr); ok {
				if expr.UnName {
					continue
				}
				if idx, ok := ident[expr.Name]; ok {
					e := NewExportMeta(expr.Name, expr, source, idx)
					exports2[expr.Name] = e
					delete(ident, expr.Name)
				}
			}
		case *ast.StructDeclStmt:
			exportStruct[stmt.Name.Value] = stmt
			delete(ident, stmt.Name.Value)
		}
	}

	// 注册
	er[source] = exports2

	importStruct := map[string]bool{}

	// 处理exportStruct, 收集组合了 import中的
	for _, stru := range exportStruct {
		fields := stru.Fields
		for _, fields := range fields {
			fname := fields.Name.Value
			if fields.IsEmbed { // 组合，在本模块中找，没有则是在import中，进行记录
				if _, ok := structs[fname]; !ok {
					importStruct[fname] = true
				}
			}
		}
	}

	importStructMeta := map[string]*ExportMeta{}
	// 还有，从imports中找
	// 1: ident 10:importStruct
	if len(ident) > 0 || len(importStruct) > 0 {
		kind := 0
		for _, imp := range imports {
			sr := ResolveImportPath(source, imp.Source)
			for _, spe := range imp.Specifiers {
				// ident
				idx, ok := ident[spe.Local]
				if ok {
					kind = kind | 1
				}
				// importStruct
				_, ok = importStruct[spe.Local]
				if ok {
					kind = kind | 2
				}

				// all not
				if kind == 0 {
					continue
				}

				ep2, ok := er[sr]
				if !ok {
					ep2 = resolveExports(l, sr, er)
				}
				ep, ok := ep2[spe.Imported]
				if !ok {
					msg := fmt.Sprintf("no exported named %q in file %s", spe.Imported, sr)
					panic(msg)
				}
				nep := NewExportMeta(spe.Local, ep.Ast, ep.Source, idx)

				if (kind & 1) == 1 {
					nep.SetReExportInfo(ep.Name)
					exports2[spe.Local] = nep
				}
				if (kind & 2) == 2 {
					importStructMeta[spe.Local] = nep
				}
			}
		}
	}

	// 处理 export struct, 得到字段总数，收集直接方法
	//for name, struc := range exportStruct {
	//}

	return exports2
}
