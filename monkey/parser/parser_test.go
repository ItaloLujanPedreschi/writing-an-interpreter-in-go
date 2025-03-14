package parser

import (
  "testing"
  "monkey/ast"
  "monkey/lexer"
)

func TestLetStatements(t *testing.T) {
  input := `
let x = 5;
let y = 10;
let foobar = 838383;
`

  l := lexer.New(input)
  p := New(l)

  program := p.ParseProgram()
  checkParserErrors(t, p)
  if program == nil {
    t.Fatalf("ParseProgram() returned nil")
  }
  if len(program.Statements) != 3 {
    t.Fatalf("program.Statements does not contain 3 statements. got=%d",
      len(program.Statements))
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
    if !testLetStatement(t, stmt, tt.expectedIdentifier) {
      return
    }
  }
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
  if s.TokenLiteral() != "let" {
    t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
    return false
  }

  letStatement, ok := s.(*ast.LetStatement)
  if !ok {
    t.Errorf("s not *ast.LetStatement. got=%T", s)
    return false
  }
  
  if letStatement.Name.Value != name {
    t.Errorf("letStatement.Name.Value not '%s'. got=%s", name, letStatement.Name.Value)
    return false
  }

  if letStatement.Name.TokenLiteral() != name {
    t.Errorf("s.Name not '%s'. got=%s", name, letStatement.Name)
    return false
  }

  return true
}

func TestReturnStatements(t *testing.T) {
  input := `
return 5;
return 10;
return 993322;
`

  l := lexer.New(input)
  p := New(l)

  program := p.ParseProgram()
  checkParserErrors(t, p)

  if len(program.Statements) != 3 {
    t.Fatalf("program.Statements does not contain 3 statements. got=%d",
      len(program.Statements))
  }

  for _, statement := range program.Statements {
    returnStatement, ok := statement.(*ast.ReturnStatement)
    if !ok {
      t.Errorf("statement not *ast.returnStatement. got %T", statement)
      continue
    }
    if returnStatement.TokenLiteral() != "return" {
      t.Errorf("returnStatement.TokenLiteral not 'return', got %q",
        returnStatement.TokenLiteral())
    }
  }
}

func checkParserErrors(t *testing.T, p *Parser) {
  errors := p.Errors()
  if len(errors) == 0 {
    return
  }

  t.Errorf("parser has %d errors", len(errors))
  for _, msg := range errors {
    t.Errorf("parser error: %q", msg)
  }
  t.FailNow()
}

