package parser

import (
	"interpreter/ast"
	"interpreter/lexer"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() errored with nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements did not return 3 statements. Parsed %d statements instead.", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]

		if !testLetStatementIdentifier(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatementIdentifier(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not parsed to `let`. Parsed `%q` instead.", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. Received `%T` instead.", s)
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not `%s`. Received `%s` instead.", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not `%s`. Received `%s` instead.", name, letStmt.Name)
		return false
	}

	return true
}
