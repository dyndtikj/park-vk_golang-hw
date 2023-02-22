package parser

import (
	"fmt"
	"homework/hw1_part2/calculator/token"
)

func Parse(input string) ([]token.Token, error) {
	ter := Tokenizer{input, 0}
	var result []token.Token
	for {
		t, err := ter.NextToken()
		if err != nil {
			return []token.Token{}, fmt.Errorf("cant parse next token %w", err)
		}
		if t.Type == token.EofType {
			break
		}
		result = append(result, t)
	}
	return result, nil
}
