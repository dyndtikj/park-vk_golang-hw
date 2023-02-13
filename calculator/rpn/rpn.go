package rpn

import (
	"calculator/calculator/token"
	"fmt"
)

// Создает слайс токенов в обратной польской нотации
func createRPN(tokens []token.Token) (result []token.Token) {
	for _, t := range tokens {
		fmt.Println(t)
	}
	return
}
