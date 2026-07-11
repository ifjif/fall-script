package semantic

import "zzc/fall-script/src/ir"

func (c *Analyzer) enterFunction() {
	newSymbolTable := ir.NewEnclosedSymbolTableForFn(c.SymbolTable)
	c.SymbolTable = newSymbolTable
}

func (c *Analyzer) leaveFunction() {
	outer := c.SymbolTable.Outer()
	c.SymbolTable = outer
}

func (c *Analyzer) enterBlock() {
	newSymbolTable := ir.NewEnclosedSymbolTable(c.SymbolTable)
	newSymbolTable.SetDefNum(c.SymbolTable.DefNum())
	c.SymbolTable = newSymbolTable
}

func (c *Analyzer) leaveBlock() {
	maxNum := c.SymbolTable.MaxNum()
	outer := c.SymbolTable.Outer()
	outer.CASMaxNum(maxNum)
	c.SymbolTable = outer
}
