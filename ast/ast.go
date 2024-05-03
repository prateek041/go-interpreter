package ast

import "github.com/prateek041/go-interpreter/token"

type Node interface {
	// TokenLiteral method returns the actual literal of that token. i.e. "let", "+" etc.
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement struct holds all the information needed to represent a let statement binding.
// let x = 5 has three major things
// let <identifier> = <expression>
// let: is a token token.LET.
// identifier: is a token, which can be a simple name but also an expression itself. so we are writing re-uable logic for it. currently it only represents the name of the variable.
// value: the evaluated value of an expression.
//

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) statementNode() {}

type Identifier struct {
	Token token.Token // here the token is IDENT.
	Value string      // value is the actual name of the variable, "x", "y" etc.
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
