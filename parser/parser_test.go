package parser

import (
	"testing"

	"github.com/prateek041/go-interpreter/ast"
	"github.com/prateek041/go-interpreter/lexer"
)

func TestLetStatement(t *testing.T) {
	input := `let x = 5;
  let y = 10;
  let foobar = 123456`

	l := lexer.New(input)

	p := New(l)

	ast := p.ParseProgram()

	if ast == nil {
		t.Fatalf("Parsing program returned nil")
	}

	if len(ast.Statements) != 3 {
		t.Fatalf("ast.Statements do not contain 3 statemtns. got=%d", len(ast.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := ast.Statements[i]

		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, stmt ast.Statement, name string) {
	if stmt.TokenLiteral() != "let" {
		// error
	}

	letstmt, ok := stmt.(*ast.LetStatement)
}
