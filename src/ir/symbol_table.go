package ir

type (
	SymbolScope      string
	SymbolTableScope string
)

const (
	BLOCK    = "BLOCK"
	FUNCTION = "FUNCTION"
)

const (
	GLOBAL  = "GLOBAL"
	LOCAL   = "LOCAL"
	BUILTIN = "BUILTIN"
	FREE    = "FREE"
	FN      = "FN"
)

type Symbol struct {
	Name     string
	Pos      int
	Scope    SymbolScope
	Captured bool
}

type SymbolTable struct {
	outer       *SymbolTable
	store       map[string]*Symbol
	globalNum   int
	defNum      int
	maxNum      int
	FreeSymbols []*Symbol
	scope       SymbolTableScope
}

func (st *SymbolTable) Outer() *SymbolTable {
	return st.outer
}

func (st *SymbolTable) DefNum() int {
	return st.defNum
}

func (st *SymbolTable) SetDefNum(defNum int) {
	st.defNum = defNum
}

func (st *SymbolTable) GlobalNum() int {
	return st.globalNum
}

func (st *SymbolTable) MaxNum() int {
	return st.maxNum
}

func (st *SymbolTable) CASMaxNum(num int) {
	if st.maxNum < num {
		st.maxNum = num
	}
}

func NewSymbolTable() *SymbolTable {
	store := make(map[string]*Symbol)
	fs := []*Symbol{}
	st := &SymbolTable{store: store, FreeSymbols: fs, scope: BLOCK}
	return st
}

func NewEnclosedSymbolTable(outer *SymbolTable) *SymbolTable {
	newSt := NewSymbolTable()
	newSt.outer = outer
	return newSt
}

func NewEnclosedSymbolTableForFn(outer *SymbolTable) *SymbolTable {
	newSt := NewSymbolTable()
	newSt.outer = outer
	newSt.scope = FUNCTION
	return newSt
}

func (st *SymbolTable) Define(name string) *Symbol {
	symbol := &Symbol{Name: name, Pos: st.defNum}

	if st.outer.outer == nil {
		symbol.Scope = GLOBAL
		symbol.Pos = st.globalNum
		st.globalNum++
	} else {
		symbol.Scope = LOCAL
		st.defNum++
		if st.maxNum < st.defNum {
			st.maxNum = st.defNum
		}
	}

	st.store[name] = symbol
	return symbol
}

func (st *SymbolTable) ResolveSelf(name string) (*Symbol, bool) {
	sym, ok := st.store[name]
	return sym, ok
}

func (st *SymbolTable) Resolve(name string) (*Symbol, bool) {
	sym, ok := st.store[name]

	if !ok && st.outer != nil {
		sym, ok = st.outer.Resolve(name)

		if !ok {
			return sym, ok
		}

		if sym.Scope == GLOBAL || sym.Scope == BUILTIN || st.scope == BLOCK {
			return sym, ok
		}

		free := st.defineFree(sym)
		return free, true
	}

	return sym, ok
}

func (st *SymbolTable) defineFree(origin *Symbol) *Symbol {
	if origin.Scope == LOCAL && !origin.Captured {
		origin.Captured = true
	}
	st.FreeSymbols = append(st.FreeSymbols, origin)

	free := &Symbol{Name: origin.Name, Pos: len(st.FreeSymbols) - 1, Scope: FREE}
	st.store[origin.Name] = free

	return free
}

func (st *SymbolTable) DefineBuiltin(index int, name string) *Symbol {
	sym := &Symbol{Name: name, Pos: index, Scope: BUILTIN}
	st.store[name] = sym
	return sym
}

func (st *SymbolTable) DefineFunction(name string) *Symbol {
	sym := &Symbol{Name: name, Pos: 0, Scope: FN}
	st.store[name] = sym
	return sym
}
