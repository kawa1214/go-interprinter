package token

// Type トークンのタイプ．雑多なコードやヘルパー関数を用意せず簡単にデバッグをするため，string型としている．
type Type string

const (
	// ILLEGAL トークンや文字が未知
	ILLEGAL = "ILLEGAL"
	// EOF ファイルの終端
	EOF = "EOF"

	// IDENT is a
	IDENT = "IDENT" // add, foobar, x, y, ...
	// INT is a
	INT = "INT" // 1343456

	// ASSIGN is a
	ASSIGN = "="
	// PLUS is a
	PLUS = "+"
	// MINUS is a
	MINUS = "-"
	// BANG is a
	BANG = "!"
	// ASTERISK is a
	ASTERISK = "*"
	// SLASH is a
	SLASH = "/"

	// LT is a
	LT = "<"
	// GT is a
	GT = ">"

	// EQ is a
	EQ = "=="
	// NOTEQ is a
	NOTEQ = "!="

	// COMMA is a
	COMMA = ","
	// SEMICOLON is a
	SEMICOLON = ";"

	// LPAREN is a
	LPAREN = "("
	// RPAREN is a
	RPAREN = ")"
	// LBRACE is a
	LBRACE = "{"
	// RBRACE is a
	RBRACE = "}"

	// FUNCTION is a
	FUNCTION = "FUNCTION"
	// LET is a
	LET = "LET"
	// TRUE is a
	TRUE = "TRUE"
	// FALSE is a
	FALSE = "FALSE"
	// IF is a
	IF = "IF"
	// ELSE is a
	ELSE = "ELSE"
	// RETURN is a
	RETURN = "RETURN"
)

// Token is
type Token struct {
	Type    Type
	Literal string
}

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent is check indent
func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
