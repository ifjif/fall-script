package parser

import "zzc/fall-script/src/token"

type Precedence int

const (
	_ Precedence = iota
	LOWEST
	ASSIGN_P    // =
	LOGIC       // || &&
	EQUALS      // == !=
	LESSGREATER // > >= < <=
	SUM         // +, -
	PRODUCT     // *, /
	PREFIX      // - !
	CALL        // fn()
	INDEX       // array[index]
	MEMBER      // a.b
)

var precedences = map[token.TokenType]Precedence{
	token.ASSIGN:   ASSIGN_P,
	token.AND:      LOGIC,
	token.OR:       LOGIC,
	token.EQ:       EQUALS,
	token.NEQ:      EQUALS,
	token.LT:       LESSGREATER,
	token.LE:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.GE:       LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.ASTERISK: PRODUCT,
	token.SLASH:    PRODUCT,
	token.LPAREN:   CALL,
	token.LBRACKET: INDEX,
	token.DOT:      MEMBER,
}

func (p *Parser) curPrecedence() Precedence {
	if p, ok := precedences[p.curType()]; ok {
		return p
	}

	return LOWEST
}

func (p *Parser) peekPrecedence() Precedence {
	if p, ok := precedences[p.nexType()]; ok {
		return p
	}

	return LOWEST
}
