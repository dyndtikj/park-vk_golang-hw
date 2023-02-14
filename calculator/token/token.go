package token

// Псевдоним для типа токена
type Type int

// все возможные типы для токенов
const (
	EOF = iota
	NUMBER
	OPERATOR
	L_PAR
	R_PAR
)

const (
	//EOF_LIT    = ''
	PLUS_LIT   = '+'
	MINUS_LIT  = '-'
	MUL_LIT    = '*'
	DIV_LIT    = '/'
	L_PAR_LIT  = '('
	R_PART_LIT = ')'
)

var Priority = map[byte]int{
	PLUS_LIT:  1,
	MINUS_LIT: 1,
	MUL_LIT:   2,
	DIV_LIT:   2,
}

type Token struct {
	Type    Type
	Literal string
}

func NewToken(t Type, l string) Token {
	return Token{
		Type:    t,
		Literal: l,
	}
}
