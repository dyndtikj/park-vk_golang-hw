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

// Надо немного подтюнить чтоб так явно не работать с литералами
var Priority = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
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
