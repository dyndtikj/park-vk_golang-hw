package calc_util

type TokenType int

// language
const (
	EOF = iota
	NUMBER
	OPERATOR
	LPAR
	RPAR
)

// Надо немного подтюнить чтоб так явно не работать с литералами
var Priority = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(t TokenType, l string) Token {
	return Token{
		Type:    t,
		Literal: l,
	}
}
