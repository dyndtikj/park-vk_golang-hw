package parser

import (
	"calculator/calculator/token"
	"errors"
)

type Tokenizer struct {
	input  string
	cursor int
}

func New(input string) *Tokenizer {
	return &Tokenizer{
		input:  input,
		cursor: 0,
	}
}

func (t *Tokenizer) NextToken() (token.Token, error) {
	t.skipSpace()

	switch t.peekChar() {
	case token.PlusLit, token.MinusLit, token.MulLit, token.DivLit:
		tok := token.NewToken(token.OpType, string(t.peekChar()))
		t.cursor++
		return tok, nil
	case token.LParLit:
		tok := token.NewToken(token.LparType, string(t.peekChar()))
		t.cursor++
		return tok, nil
	case token.RPartLit:
		tok := token.NewToken(token.RparType, string(t.peekChar()))
		t.cursor++
		return tok, nil

	case 0:
		return token.NewToken(token.EofType, ""), nil

	default:
		if t.isDigit(t.peekChar()) {
			var number string
			for t.isDigit(t.peekChar()) {
				number += string(t.peekChar())
				t.cursor++
			}

			return token.NewToken(token.NumType, number), nil
		}
		return token.Token{}, errors.New("Can't find token from moment: " + t.input[t.cursor:])
	}
}

func (t *Tokenizer) peekChar() byte {
	// дошли до конца
	if t.cursor >= len(t.input) {
		return 0
	}
	return t.input[t.cursor]
}

func (t *Tokenizer) skipSpace() {
	for t.isSpace(t.peekChar()) {
		t.cursor++
	}
}

func (t *Tokenizer) isSpace(ch byte) bool {
	switch ch {
	case '\n', ' ', '\t':
		return true
	}
	return false
}

func (t *Tokenizer) isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}
