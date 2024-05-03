package lexer

import (
	"fmt"
	"io"

	"github.com/prateek041/go-interpreter/token"
)

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
	// eat any whitespace character from the input.
	l.eatWhiteSpace()

	switch l.ch {
	case '=':
		if l.peakChar() == '=' {
			firstChar := l.ch
			l.readChar()
			literal := string(firstChar) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = token.Token{Type: token.ASSIGN, Literal: string(l.ch)}
		}

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

	case '-':
		tok = token.Token{Type: token.MINUS, Literal: string(l.ch)}

	case '!':
		if l.peakChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = token.Token{Type: token.BANG, Literal: string(l.ch)}
		}

	case '*':
		tok = token.Token{Type: token.ASTERISK, Literal: string(l.ch)}

	case '/':
		tok = token.Token{Type: token.SLASH, Literal: string(l.ch)}

	case '<':
		tok = token.Token{Type: token.LT, Literal: string(l.ch)}

	case '>':
		tok = token.Token{Type: token.GT, Literal: string(l.ch)}

	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	default:
		if isLetter(l.ch) {
			identifier := l.getIdentifier()
			tok.Literal = identifier
			tok.Type = token.GetIdent(identifier)
			return tok
		} else if isDigit(l.ch) {
			number := l.getNumber()
			fmt.Println("number is", number)
			tok.Literal = number
			tok.Type = token.INT
			return tok
		} else {
			tok = token.Token{Type: token.ILLEGAL, Literal: string(l.ch)}
		}

	}

	l.readChar()
	return tok
}

func (l *Lexer) SpitTokens(out io.Writer) {
	for l.position <= len(l.input) {
		tok := l.NextToken()
		fmt.Fprintf(out, "%+v\n", tok)
	}
}

func (l *Lexer) eatWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) getIdentifier() string {
	startPoint := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	tok := l.input[startPoint:l.position]
	return tok
}

func (l *Lexer) getNumber() string {
	startPoint := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[startPoint:l.position]
}

func (l *Lexer) peakChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	ch := l.input[l.readPosition]
	return ch
}

func isDigit(ch byte) bool {
	if ch >= '0' && ch <= '9' {
		return true
	}
	return false
}

func isLetter(ch byte) bool {
	if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_' {
		return true
	}
	return false
}
