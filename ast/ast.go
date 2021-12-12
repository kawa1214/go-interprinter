package ast

import (
	"github.com/kawa1214/go-interprinter/token"
)

// 構文解析器(パーサ) 入力をデータ構造に変換
// ソースコードの内部表現としては、構文木もしくは抽象構文木が用いられる。

// ASTの全てのノードはNodeインターフェイスを実装必要がある。TokenLiteral()メソッドを提供しなければならない。
// TokenLiteralメソッドは、ノードが関連付けられているトークンのリテラル値を返す。
// TokenLiteral()はデバッグとテストだけに用い、ASTは互いに接続されたNodeだけで構成される。

type Node interface {
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

// ASTはのルートノードになる
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

// Expressionはノードの種類を
type LetStatement struct {
	Token token.Token // token.LET トークン
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token // token.IDENT トークン
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

type ReturnStatement struct {
	Token       token.Token // 'return' トークン
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
