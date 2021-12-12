package parser

import (
	"fmt"

	"github.com/kawa1214/go-interprinter/ast"
	"github.com/kawa1214/go-interprinter/lexer"
	"github.com/kawa1214/go-interprinter/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
  errors []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
    l: l,
    errors: []string{},
  }

	// 2つのトークンを読み込む。curTokenとpeekTokenがセットされる
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

  // ASTのルートノードを生成し、token.EOFに到達するまで、入力のトークンを繰り返し呼ぶ。
	for p.curToken.Type != token.EOF {
    stmt := p.parseStatement()
    if stmt != nil {
      program.Statements = append(program.Statements, stmt)
    }
    p.nextToken()
	}
  return program
}

func (p *Parser) parseStatement() ast.Statement {
  switch p.curToken.Type {
  case token.LET:
    return p.parseLetStatement()
  default:
    return nil
  }
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
  stmt := &ast.LetStatement{Token: p.curToken}
  
  if !p.expectPeek(token.IDENT) {
    return nil
  }

  stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

  if !p.expectPeek(token.ASSIGN) {
    return nil
  }

  for !p.curTokenIs(token.SEMICOLON) {
    p.nextToken()
  }

  return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
  return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
  return p.peekToken.Type == t
}

// アサーション関数：peekTokenの型をチェクし、正しければnextTokenを呼び、トークンを進める。
func (p *Parser) expectPeek(t token.TokenType) bool {
  if p.peekTokenIs(t) {
    p.nextToken()
    return true
  } else {
    p.peekError(t)
    return false
  }
}

func (p *Parser) Errors() []string {
  return p.errors;
}

func (p *Parser) peekError(t token.TokenType){
  msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
  p.errors = append(p.errors, msg)
}