package semantic

import (
	"fmt"

	"zzc/fall-script/src/ast"
	"zzc/fall-script/src/ir"
	"zzc/fall-script/src/object"
)

type FieldInfo struct {
	Name     string
	IsEmbed  bool
	Index    int
	Children map[string]*FieldInfo
}

func (a *Analyzer) analyzeStruct(structDeclStmts map[string]*ast.StructDeclStmt, methodDeclExprs map[string][]*ast.MethodDeclExpr) []ir.Stmt {
	ses := []ir.Stmt{}

	for sname, structDeclStmt := range structDeclStmts {
		sym, ok := a.SymbolTable.Resolve(sname)
		if !ok {
			sym = a.SymbolTable.Define(sname)
		}
		se := &ir.StructDeclStmt{
			Name: sym,
		}

		count, fis := CalcStructFields(structDeclStmts, structDeclStmt, 0)

		fmt.Printf("%s 的字段(%d):\n", sname, count)
		printFields(fis)

		// 处理自己的方法
		methods := methodDeclExprs[sname]
		mis := make(map[string]*object.MethodRef)
		for _, method := range methods {
			mname := method.Fn.Name
			sym := a.SymbolTable.Define(mname)
			mis[mname] = &object.MethodRef{
				TargetStructMeta: -1,
				Name:             mname,
				Index:            sym.Pos,
			}
		}

		// 进行 一级平铺, 它自己级 一级组合
		nfis := make(map[string]*object.FieldInfo)

		for name, f := range fis {
			ff := &object.FieldInfo{
				Name:    f.Name,
				Index:   f.Index,
				IsEmbed: f.IsEmbed,
			}
			if f.IsEmbed {
				sym, ok := a.SymbolTable.Resolve(name)
				if !ok {
					sym = a.SymbolTable.Define(name)
				}
				ff.TargetStructMeta = sym.Pos
				// 方法
				methods := methodDeclExprs[f.Name]
				for _, method := range methods {
					mname := method.Fn.Name
					// 不存在，进行promoted
					if _, ok := mis[mname]; !ok {
						mis[mname] = &object.MethodRef{
							TargetStructMeta: sym.Pos,
							Name:             mname,
							Index:            -1,
						}
					}
				}
			}
			nfis[name] = ff
		}

		fmt.Println()

		metaData := &object.StructMeta{
			Name:       sname,
			FieldCount: count,
			Fields:     nfis,
			Methods:    mis,
		}

		se.MetaData = metaData
		ses = append(ses, se)
	}

	return ses
}

func CalcStructFields(asts map[string]*ast.StructDeclStmt, structAst *ast.StructDeclStmt, offset int) (int, map[string]*FieldInfo) {
	fis := make(map[string]*FieldInfo)
	fields := structAst.Fields
	for _, field := range fields {
		fi := &FieldInfo{
			Name:    field.Name.Value,
			Index:   offset,
			IsEmbed: field.IsEmbed,
		}

		fis[fi.Name] = fi

		if field.IsEmbed {
			st := asts[field.Name.Value]
			no, ns := CalcStructFields(asts, st, offset)
			offset = no
			fi.Children = ns
		} else {
			offset++
		}

	}

	return offset, fis
}

func (a *Analyzer) analyzeMethods(methods map[string][]*ast.MethodDeclExpr) []ir.Stmt {
	fns := make([]ir.Stmt, 0)
	for _, method := range methods {
		for _, m := range method {
			fn := a.analyzeFnExpr(m.Fn)
			fns = append(fns, &ir.ExprStmt{Expr: fn})
		}
	}

	return fns
}

func printFields(fields map[string]*FieldInfo) {
	for _, f := range fields {
		fmt.Printf("name: %s index: %d isEmbed: %t\n", f.Name, f.Index, f.IsEmbed)
		if f.IsEmbed {
			for _, ff := range f.Children {
				fmt.Printf("		name: %s index: %d isEmbed: %t\n", ff.Name, ff.Index, ff.IsEmbed)
			}
		}
	}
}
