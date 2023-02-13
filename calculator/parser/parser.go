package parser

import (
	"calculator/calculator/token"
	"calculator/calculator/tokenizer"
)

// Возвращает слайс токенов
func Parse(input string) (result []token.Token, err error) {
	ter := tokenizer.New(input)

	for {
		var t token.Token
		t, err = ter.NextToken()
		if err != nil {
			return
		}
		if t.Type == token.EOF {
			break
		}
		result = append(result, t)
	}
	return
}
