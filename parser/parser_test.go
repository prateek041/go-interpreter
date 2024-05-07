package parser

import (
	"fmt"
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

	fmt.Println("Parsing started ")
	ast := p.ParseProgram()
	checkParseError(t, p)

	fmt.Println("Parsing done successfully")

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
		fmt.Println("checking this statement", stmt)

		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func checkParseError(t *testing.T, p *Parser) {
	errors := p.GetErrors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func testLetStatement(t *testing.T, stmt ast.Statement, name string) bool {
	if stmt.TokenLiteral() != "let" {
		t.Errorf("stmt.TokenLiteral not 'let'. got=%q", stmt.TokenLiteral())
		return false
	}

	// asserting if stmt, the interface, holds a concrete type LetStatement
	letstmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("stmt not *ast.LetStatement. got=%T", stmt)
		return false
	}

	if letstmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letstmt.Name.Value)
		return false
	}

	if letstmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s", name, letstmt.Name.TokenLiteral())
		return false
	}
	return true
}
