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

type action func(a, b float64) float64

var Priority = map[byte]int{
	PlusLit:  1,
	MinusLit: 1,
	MulLit:   2,
	DivLit:   2,
}

var Actions = map[byte]action{
	PlusLit: func(a, b float64) float64 {
		return a + b
	},
	MinusLit: func(a, b float64) float64 {
		return a - b
	},
	MulLit: func(a, b float64) float64 {
		return a * b
	},
	DivLit: func(a, b float64) float64 {
		return a / b
	},
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
