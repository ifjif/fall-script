package compiler

import (
	"fmt"

	"zzc/fall-script/src/ast"
	"zzc/fall-script/src/code"
	"zzc/fall-script/src/object"
)

type StructMeta2 struct {
	Sm         *object.StructMeta
	Composites []*FieldInfo
}

type FieldInfo struct {
	Name     string
	IsEmbed  bool
	Index    int
	Children map[string]*FieldInfo
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

func printFields2(fields map[string]*object.FieldInfo) {
	for _, f := range fields {
		fmt.Printf("name: %s index: %d isEmbed: %t targetMeta: %v\n", f.Name, f.Index, f.IsEmbed, f.TargetStructMeta)
	}
}

func (c *Compiler) resolveMethods(struct2s map[string]*StructMeta2) {
	c.resolveMethods2(c.methods)
	// 处理组合
	for _, structM := range struct2s {
		compo := structM.Composites
		methods := structM.Sm.Methods
		for _, field := range compo {
			cn := field.Name
			sym, ok := c.SymbolTable.Resolve(cn)
			if !ok {
				panic("not exist field")
			}
			cs := c.structs[cn]

			for name, method := range cs.Methods {
				// 存在忽略
				if _, ok := methods[name]; !ok {
					methods[name] = &object.MethodRef{
						Name:             method.Name,
						Index:            method.Index,
						TargetStructMeta: sym.Pos,
					}
				}
			}
		}
	}
}

func (c *Compiler) resolveMethods2(methodsAst map[string][]*ast.MethodDeclExpr) {
	for structName, methods := range methodsAst {
		mrs := c.compileMethods(methods)
		stru := c.structs[structName]
		stru.Methods = mrs
	}
}

func (c *Compiler) compileMethods(methods []*ast.MethodDeclExpr) map[string]*object.MethodRef {
	mrs := map[string]*object.MethodRef{}
	for _, method := range methods {
		name := method.Fn.Ident.Value
		c.compileFnExpr(method.Fn)

		mr := &object.MethodRef{
			TargetStructMeta: -1,
			Name:             name,
			Index:            c.SymbolTable.globalNum - 1,
		}
		mrs[name] = mr
	}

	return mrs
}

func (c *Compiler) resolveStructs() {
	struct2s, structs := resolveStructs2(c.structAsts)
	c.structs = structs
	c.compileStructs(struct2s)
}

func (c *Compiler) compileStructs(struct2s map[string]*StructMeta2) {
	for name, sm := range c.structs {
		sym := c.SymbolTable.Define(name)
		c.emit(code.Const, c.addConstant(sm))
		c.storeSymbol(sym)
	}

	c.resolveTargetMeta4EmbedField(struct2s)
	c.resolveMethods(struct2s)
}

// 为所有组合字段设置 targetMetaIndex
func (c *Compiler) resolveTargetMeta4EmbedField(struct2s map[string]*StructMeta2) {
	for _, structM := range struct2s {
		stru := structM.Sm
		fields := structM.Composites
		for _, field := range fields {
			fn := field.Name
			sym, ok := c.SymbolTable.Resolve(fn)
			if !ok {
				panic("struct symbol not exist")
			}
			stru.Fields[fn].TargetStructMeta = sym.Pos
		}
	}
}

func resolveStructs2(asts map[string]*ast.StructDeclStmt) (map[string]*StructMeta2, map[string]*object.StructMeta) {
	structMeta2s := map[string]*StructMeta2{}
	cstrus := map[string]*object.StructMeta{}
	for sname, ast := range asts {
		count, fis := doResolveStruct2(asts, ast, 0)

		name := ast.Name
		fmt.Printf("%s 的字段(%d):\n", name, count)
		printFields(fis)

		fmt.Println("一级平铺后")

		// 进行 一级平铺, 它自己级 一级组合
		nfis := make(map[string]*object.FieldInfo)

		comps := []*FieldInfo{}
		for name, f := range fis {
			nfis[name] = &object.FieldInfo{
				Name:    f.Name,
				Index:   f.Index,
				IsEmbed: f.IsEmbed,
			}
			if f.IsEmbed {
				comps = append(comps, f)
			}
		}

		// 处理一级组合
		for _, f := range comps {
			for n2, f2 := range f.Children {
				if _, ok := nfis[n2]; !ok {
					nfis[n2] = &object.FieldInfo{
						Name:    f2.Name,
						Index:   f2.Index,
						IsEmbed: f2.IsEmbed,
					}
				}
			}
		}
		fmt.Println()

		cstru := &object.StructMeta{
			Name:       sname,
			FieldCount: count,
			Fields:     nfis,
		}
		cstrus[sname] = cstru
		structMeta2s[sname] = &StructMeta2{
			Sm:         cstru,
			Composites: comps,
		}
	}

	fmt.Println("struct meta:")
	for name, cstru := range structMeta2s {
		fmt.Printf("name: %s\n", name)
		fmt.Printf("field count: %d\n", cstru.Sm.FieldCount)
		fmt.Println("fields: ")
		printFields2(cstru.Sm.Fields)
		fmt.Println()
	}

	return structMeta2s, cstrus
}

func doResolveStruct2(asts map[string]*ast.StructDeclStmt, structAst *ast.StructDeclStmt, offset int) (int, map[string]*FieldInfo) {
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
			no, ns := doResolveStruct2(asts, st, offset)
			offset = no
			fi.Children = ns
		} else {
			offset++
		}

	}

	return offset, fis
}
