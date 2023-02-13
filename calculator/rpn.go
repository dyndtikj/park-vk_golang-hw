package calculator

import (
	"calculator/calculator/token"
	"fmt"
)

func createRPN(tokens []token.Token) (result []token.Token) {
	for _, t := range tokens {
		fmt.Println(t)
	}
	return
}
