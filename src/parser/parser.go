package parser

import (
	. "zzc/fall-script/src/ast"
	"zzc/fall-script/src/lexer"
	. "zzc/fall-script/src/token"
)

type Parser struct {
	l              *lexer.Lexer
	curToken       Token
	nexToken       Token
	errors         []string
	prefixParseFns map[TokenType]prefixParseFn
	infixParseFns  map[TokenType]infixParseFn
}

func NewParser(input string) *Parser {
	l := lexer.NewLexer(input)
	p := &Parser{
		l:      l,
		errors: []string{},
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
	return program
}

func (p *Parser) Errors() []string {
	return p.errors
}
