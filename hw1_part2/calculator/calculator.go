package calculator

import (
	"errors"
	"fmt"
	"homework/hw1_part2/calculator/parser"
	"homework/hw1_part2/calculator/rpn"
	"homework/hw1_part2/calculator/token"
)

var (
	ErrWrongParentheses = errors.New("wrong parentheses")
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
			return ErrWrongParentheses
		}
	}
	if counter != 0 {
		return ErrWrongParentheses
	}
	return nil
}

func Calculate(input string) (float64, error) {
	if err := checkExpression(input); err != nil {
		return 0, fmt.Errorf("failed to check expression %w", err)
	}

	tokens, err := parser.Parse(input)
	if err != nil {
		return 0, fmt.Errorf("failed to parse expression %w", err)
	}

	rpnTokens, err := rpn.CreateRPN(tokens)
	if err != nil {
		return 0, fmt.Errorf("failed to create RPN %w", err)
	}
	result, err := rpn.EvaluateRpn(rpnTokens)
	if err != nil {
		return 0, fmt.Errorf("failed to evaluste RPN %w", err)
	}
	return result, nil
}
