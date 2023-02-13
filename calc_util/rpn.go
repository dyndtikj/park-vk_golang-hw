package calc_util

import "fmt"

func createRPN(tokens []Token) (result []Token) {
	for _, token := range tokens {
		fmt.Println(token)
	}
	return
}
