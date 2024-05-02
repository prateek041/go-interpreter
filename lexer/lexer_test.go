package lexer

import (
	"fmt"
	"testing"

	"github.com/prateek041/go-interpreter/token"
)

func TestNextFunction(t *testing.T) {
	input := "=+(){},;"
	fmt.Println("This test was run")

	tests := []token.Token{
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.PLUS, Literal: "+"},
		{Type: token.LPAREN, Literal: "("},
		{Type: token.RPAREN, Literal: ")"},
		{Type: token.LBRACE, Literal: "{"},
		{Type: token.RBRACE, Literal: "}"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.SEMICOLON, Literal: ";"},
	}

	l := New(input)

	for i, tt := range tests {
		// we expect a Token struct here. {Type: token.ASSIGN, Literal: "="}
		tok := l.NextToken()

		if tok.Type != tt.Type {
			t.Fatalf("test[%d] - token type wrong, expected=%q, got=%q", i, tt.Type, tok.Type)
		}

		if tok.Literal != tt.Literal {
			t.Fatalf("test[%d] - token Literal wrong, expected=%q, got=%q", i, tt.Literal, tok.Literal)
		}
	}
}
