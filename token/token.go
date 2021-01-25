package token

// TokenType - many different values as tokentypes
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	// identifiers and literals
	IDENT = "IDENT"
	INT   = "INT"
	// operators
	ASSIGN = "="
	PLUS   = "+"

	// delimiter
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
