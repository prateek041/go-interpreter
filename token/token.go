package token

// this will tokens supported by our language.

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals
	IDENT = "IDENT"

	// operators =, +
	ASSIGN = "="
	PLUS   = "+"

	// delimeters
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	COMMA     = ","
	SEMICOLON = ";"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type Token struct {
	// TODO: Think about why a different Token Type was defined instead of just using a string.
	Type    string
	Literal string
}
