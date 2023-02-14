package parser

import (
	"calculator/calculator/token"
)

func Parse(input string) (result []token.Token, err error) {
	ter := New(input)

	for {
		var t token.Token
		t, err = ter.NextToken()
		if err != nil {
			return
		}
		if t.Type == token.EofType {
			break
		}
		result = append(result, t)
	}
	return
}
