package ast

import "github.com/prateek041/go-interpreter/token"

type Node interface {
	// TokenLiteral method returns the actual literal of that token. i.e. "let", "+" etc.
	TokenLiteral() string
}

// Statement is a node in the AST, which does not generate a value
type Statement interface {
	Node
	statementNode()
}

// Expression is a node in AST that generates a value.
type Expression interface {
	Node
	expressionNode()
}

// Program is a collection of statements
type Program struct {
	Statements []Statement
}

// TokenLiteral method implements the Node interface
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
// identifier: this holds the name of the binding.
// value: the evaluated value of an expression.
type LetStatement struct {
	Token token.Token
	Value *Expression
	Name  Identifier
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) statementNode() {}

// Identifier holds the name of the binding (variable).
type Identifier struct {
	Token token.Token // here the token is IDENT.
	Value string      // value is the actual name of the variable, "x", "y" etc.
}

// expressionNode methods implements the ExpressionNode interface
// Why is the identifier implmeneting expressionNode? simply because statements like let x = y, here y is an identifier
// and it is also producing values. So, to keep the number of different node types small, Identifier can have a Token
// as well as a Value.
func (i *Identifier) expressionNode() {}

// TokenLiteral methods implements the Node interface
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
