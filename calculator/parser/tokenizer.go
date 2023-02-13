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
	case '+', '-', '/', '*':
		tok := token.NewToken(token.OPERATOR, string(t.peekChar()))
		t.cursor++
		return tok, nil
	case '(':
		tok := token.NewToken(token.L_PAR, string(t.peekChar()))
		t.cursor++
		return tok, nil
	case ')':
		tok := token.NewToken(token.R_PAR, string(t.peekChar()))
		t.cursor++
		return tok, nil

	case 0:
		return token.NewToken(token.EOF, ""), nil

	default:
		if t.isDigit(t.peekChar()) {

			var number string
			for t.isDigit(t.peekChar()) {
				number += string(t.peekChar())
				t.cursor++
			}

			return token.NewToken(token.NUMBER, number), nil
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
