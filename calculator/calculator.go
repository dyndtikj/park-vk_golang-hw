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
		if c == token.L_PAR_LIT {
			counter++
		} else if c == token.R_PART_LIT {
			counter--
		}
		if counter < 0 {
			err := errors.New("wrong parentheses")
			return err
		}
	}
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

	rpnTokens, err := rpn.CreateRPN(tokens)
	if err != nil {
		return 0, err
	}
	res, err := rpn.EvaluateRpn(rpnTokens)
	if err != nil {
		return 0, err
	}
	return res, nil
}
