package parser

import (
	"calculator/calculator/token"
	"errors"
)

type Tokenizer struct {
	input  string
	cursor int
}

func (t *Tokenizer) NextToken() (token.Token, error) {
	t.skipSpace()

	switch t.peekChar() {
	case token.PlusLit, token.MinusLit, token.MulLit, token.DivLit:
		tok := token.Token{Type: token.OpType, Literal: string(t.peekChar())}
		t.cursor++
		return tok, nil
	case token.LParLit:
		tok := token.Token{Type: token.LparType, Literal: string(t.peekChar())}
		t.cursor++
		return tok, nil
	case token.RPartLit:
		tok := token.Token{Type: token.RparType, Literal: string(t.peekChar())}
		t.cursor++
		return tok, nil

	case 0:
		return token.Token{Type: token.EofType}, nil

	default:
		if t.isDigit(t.peekChar()) {
			var number string
			for t.isDigit(t.peekChar()) {
				number += string(t.peekChar())
				t.cursor++
			}

			return token.Token{Type: token.NumType, Literal: number}, nil
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
