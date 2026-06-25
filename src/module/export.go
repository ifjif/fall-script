package module

import (
	"fmt"

	"zzc/fall-script/src/ast"
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
type ExportMeta struct {
	Name      string
	Ast       ast.Node
	Source    string
	ExportIdx int
}

func NewExportMeta(name string, ast ast.Node, source string, exportIdx int) *ExportMeta {
	return &ExportMeta{
		Name:      name,
		Ast:       ast,
		ExportIdx: exportIdx,
		Source:    source,
	}
}

type ExportMetas map[string]*ExportMeta

type ExportMetasRegister map[string]ExportMetas

/*
* export 流程
* 取export中 let fn 的
* 剩下ident留存
* 从AST中 ident的
* 将 map[string]*Exports 进行注册
* 将剩下的ident从从import找
* 看源是否被注册
* 有，取出
* 没，解析
* 寻找，如果没有，panic
* 找到加入 map[string]*Export2
*
 */
func resolveExports2(l *Loader, file string, program *ast.Program, imports []*ast.ImportStmt, exports []*ast.ExportStmt, er ExportMetasRegister) ExportMetas {
	source := file

	exports2 := ExportMetas{}
	ident := map[string]int{}

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
		}
	}

	// 注册
	er[source] = exports2

	// 还有，从imports中找
	if len(ident) > 0 {
		for _, imp := range imports {
			sr := ResolveImportPath(source, imp.Source)
			for _, spe := range imp.Specifiers {
				idx, ok := ident[spe.Local]
				if !ok {
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
				nep := NewExportMeta(ep.Name, ep.Ast, ep.Source, idx)
				exports2[spe.Local] = nep
			}
		}
	}

	return exports2
}

func resolveExports(l *Loader, file string, er ExportMetasRegister) ExportMetas {
	// 加载源
	source := file
	program := l.GenerateAST(source)

	imports, exports := CollectImportsAndExports(program)

	return resolveExports2(l, file, program, imports, exports, er)
}
