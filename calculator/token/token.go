package token

// Псевдоним для типа токена
type Type int

// все возможные типы для токенов
const (
	EofType = iota
	NumType
	OpType
	LparType
	RparType
)

const (
	PlusLit  = '+'
	MinusLit = '-'
	MulLit   = '*'
	DivLit   = '/'
	LParLit  = '('
	RPartLit = ')'
)

var Priority = map[byte]int{
	PlusLit:  1,
	MinusLit: 1,
	MulLit:   2,
	DivLit:   2,
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
