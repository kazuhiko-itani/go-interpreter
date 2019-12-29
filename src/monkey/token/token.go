package token

// TokenType トークンの型
type TokenType string

// Token tokenss
type Token struct {
	Type    TokenType
	Literal string
}

// トークン定義
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT" // add, foobar, x, y
	INT   = "INT"   // 134343

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	COMMA     = ","
	SEMICOLON = ";"

	EQ 		 = "=="
	NOT_EQ = "!="

	LPAREN 	 = "("
	RPAREN 	 = ")"
	LBRACE 	 = "{"
	RBRACE 	 = "}"
	LBRACKET = "["
	RBRACKET = "]"

	STRING   = "STRING"
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE 		 = "TRUE"
	FALSE		 = "FALSE"
	IF			 = "IF"
	ELSE		 = "ELSE"
	RETURN	 = "RETURN"
)

var keywords = map[string]TokenType {
	"fn": FUNCTION,
	"let": LET,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
