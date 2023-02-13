package calculator

import (
	"calculator/calculator/parser"
	"calculator/calculator/rpn"
	"fmt"
	"log"
)

func checkExpression(input string) error {
	return nil
}

func Calculate(input string) (float64, error) {
	tokens, err := parser.Parse(input)
	if err != nil {
		log.Fatal(err)
	}

	rpnTokens, ok := rpn.CreateRPN(tokens)
	if !ok {
		fmt.Println("some error")
	}
	res, err := rpn.EvaluateRpn(rpnTokens)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}
