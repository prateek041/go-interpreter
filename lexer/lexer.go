package lexer

import "github.com/prateek041/go-interpreter/token"

type Lexer struct {
	input        string
	ch           byte
	position     int
	readPosition int
}

// initialze a new lexer with the input string and read in the first character from the input
func New(input string) Lexer {
	l := Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

// NextToken takes the current character in the lexer, and returns a token for it.
func (l *Lexer) NextToken() token.Token {
	tok := token.Token{}

	switch l.ch {
	case '=':
		tok = token.Token{Type: token.ASSIGN, Literal: string(l.ch)}

	case '+':
		tok = token.Token{Type: token.PLUS, Literal: string(l.ch)}

	case ';':
		tok = token.Token{Type: token.SEMICOLON, Literal: string(l.ch)}

	case '(':
		tok = token.Token{Type: token.LPAREN, Literal: string(l.ch)}

	case ')':
		tok = token.Token{Type: token.RPAREN, Literal: string(l.ch)}

	case '{':
		tok = token.Token{Type: token.LBRACE, Literal: string(l.ch)}

	case '}':
		tok = token.Token{Type: token.RBRACE, Literal: string(l.ch)}

	case ',':
		tok = token.Token{Type: token.COMMA, Literal: string(l.ch)}

	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}
