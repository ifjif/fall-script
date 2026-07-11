package semantic

import (
	"zzc/fall-script/src/ast"
	"zzc/fall-script/src/ir"
)

type Analyzer struct {
	errors []string

	SymbolTable *ir.SymbolTable
}

func NewAnalyzer(globalSymbol *ir.SymbolTable) *Analyzer {
	st := ir.NewEnclosedSymbolTable(globalSymbol)
	return &Analyzer{
		errors:      []string{},
		SymbolTable: st,
	}
}

func (a *Analyzer) Errors() []string {
	return a.errors
}

func (a *Analyzer) Analyze(program *ast.Program) *ir.Program {
	irStmts := []ir.Stmt{}
	// 处理 import
	ips := a.analyzeImports(program.Imports)
	irStmts = append(irStmts, ips...)
	// 处理struct
	sds := a.analyzeStruct(program.Structs, program.Methods)
	irStmts = append(irStmts, sds...)

	// 全局函数 提前定义, 初始还是按顺序进行
	for _, fn := range program.PromotedFns {
		a.SymbolTable.Define(fn)
	}

	fns := a.analyzeMethods(program.Methods)
	irStmts = append(irStmts, fns...)

	for _, stmt := range program.Stmts {
		irStmt := a.analyzeStmt(stmt)
		if irStmt != nil {
			irStmts = append(irStmts, irStmt)
		}
	}

	// 为 export 附加 globalId
	exps := a.analyzeExports(program.Exports)
	irStmts = append(irStmts, exps...)

	return &ir.Program{
		Stmts:   irStmts,
		Locals:  a.SymbolTable.MaxNum(),
		Globals: a.SymbolTable.GlobalNum(),
	}
}
