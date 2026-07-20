package parser

import (
	. "zzc/fall-script/src/ast"
	"zzc/fall-script/src/lexer"
	. "zzc/fall-script/src/token"
)

type Parser struct {
	l              *lexer.Lexer
	exprKind       TokenType
	curToken       Token
	nexToken       Token
	errors         []string
	prefixParseFns map[TokenType]prefixParseFn
	infixParseFns  map[TokenType]infixParseFn
	methods        map[string][]*MethodDeclExpr
	structs        map[string]*StructDeclStmt
	promotedFns    []string
}

func NewParser(input string) *Parser {
	l := lexer.NewLexer(input)
	p := &Parser{
		l:           l,
		errors:      []string{},
		methods:     map[string][]*MethodDeclExpr{},
		structs:     map[string]*StructDeclStmt{},
		promotedFns: []string{},
	}

	p.registeExprFn()

	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Parse() *Program {
	program := &Program{}
	program.Stmts = []StmtNode{}

	for p.curType() != EOF {
		stmt := p.parseStmt()
		if stmt != nil {
			program.Stmts = append(program.Stmts, stmt)
		}
		p.nextToken()
	}
	program.Structs = p.structs
	program.Methods = p.methods
	program.PromotedFns = p.promotedFns
	return program
}

func (p *Parser) Errors() []string {
	return p.errors
}
