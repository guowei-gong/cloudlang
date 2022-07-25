package token

type TokenType string

const (
	// ILLEGAL 未知词法单元或字符
	ILLEGAL = "ILLEGAL"
	// EOF End of file 文件结尾
	EOF = "EOF"

	// IDENT 标识符
	IDENT = "IDENT"
	// INT 字面量
	INT = "INT"

	// ASSIGN PLUS 运算符
	ASSIGN = "="
	PLUS   = "+"

	// COMMA SEMICOLON LPAREN RPAREN LBRACE RBRACE 分隔符
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// FUNCTION LET 关键字
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

// Token 词法单元
type Token struct {
	Type    TokenType // 词法单元类型
	Literal string    // 词法单元字面量
}
