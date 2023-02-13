package calculator

import (
	"calculator/calculator/parser"
	"calculator/calculator/rpn"
	"errors"
)

func checkExpression(input string) error {
	return nil
}

func Calculate(input string) (float64, error) {
	err := checkExpression(input)
	if err != nil {
		return 0, err
	}

	tokens, err := parser.Parse(input)
	if err != nil {
		return 0, err
	}

	rpnTokens, ok := rpn.CreateRPN(tokens)
	if !ok {
		return 0, errors.New("cant create RPN from your input")
	}
	res, err := rpn.EvaluateRpn(rpnTokens)
	if err != nil {
		return 0, err
	}
	return res, nil
}
