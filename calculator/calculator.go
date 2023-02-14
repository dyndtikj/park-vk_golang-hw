package calculator

import (
	"calculator/calculator/parser"
	"calculator/calculator/rpn"
	"calculator/calculator/token"
	"errors"
)

func checkExpression(input string) error {
	counter := 0
	for _, c := range input {
		if c == token.LParLit {
			counter++
		} else if c == token.RPartLit {
			counter--
		}
		if counter < 0 {
			return errors.New("wrong parentheses")
		}
	}
	if counter != 0 {
		return errors.New("wrong parentheses")
	}
	return nil
}

func Calculate(input string) (result float64, err error) {
	err = checkExpression(input)
	if err != nil {
		return
	}

	tokens, err := parser.Parse(input)
	if err != nil {
		return
	}

	rpnTokens, err := rpn.CreateRPN(tokens)
	if err != nil {
		return
	}
	result, err = rpn.EvaluateRpn(rpnTokens)
	if err != nil {
		return 0, err
	}
	return
}
