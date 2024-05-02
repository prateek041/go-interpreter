package token

// this will tokens supported by our language.

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals
	IDENT = "IDENT"
	INT   = "INT"

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
	Type    string // TODO: figure this out
	Literal string // how the character looks like in code.
}

var keywords = map[string]string{
	"let": LET,
	"fn":  FUNCTION,
}

// GetIdent takes the input unkown string, which is recognized by compiler and returns the type of it, i.e. if it is a keyword or an identifier
func GetIdent(ident string) string {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
