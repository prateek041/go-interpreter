package parser

import (
	"github.com/prateek041/go-interpreter/ast"
	"github.com/prateek041/go-interpreter/lexer"
	"github.com/prateek041/go-interpreter/token"
)

type Parser struct {
	l         *lexer.Lexer
	currToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram is a method in the parser that returns the parsed ast tree,
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
