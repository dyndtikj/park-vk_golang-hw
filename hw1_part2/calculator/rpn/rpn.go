package rpn

import (
	"fmt"
	"homework/hw1_part2/calculator/token"
	"homework/hw1_part2/stack"
	"strconv"
)

var (
	ErrRPar          = fmt.Errorf("right parentheses not closed")
	ErrNotEnoughArgs = fmt.Errorf("not enouth arguments for operator")
	ErrNoResult      = fmt.Errorf("no token with result")
)

// Создает слайс токенов в обратной польской нотации
func CreateRPN(tokens []token.Token) ([]token.Token, error) {
	st := stack.New[token.Token]()
	var postfixTokens []token.Token
	for _, t := range tokens {
		switch t.Type {
		case token.NumType:
			postfixTokens = append(postfixTokens, t)
		case token.LparType:
			st.Push(t)
		case token.RparType:
			for !st.IsEmpty() {
				topToken, ok := st.Peek()
				if !ok {
					return []token.Token{}, ErrRPar
				}
				if topToken.Type == token.LparType {
					break
				}
				postfixTokens = append(postfixTokens, topToken)
				_, _ = st.Pop()
			}
			topToken, ok := st.Peek()
			if !ok {
				return []token.Token{}, ErrRPar
			}
			if topToken.Type == token.LparType {
				// ignore err cause handled higher
				_, _ = st.Pop()
			}
		case token.OpType:
			for !st.IsEmpty() {
				topToken, ok := st.Peek()
				if !ok {
					return []token.Token{}, ErrNotEnoughArgs
				}
				if token.Priority[t.Literal[0]] > token.Priority[topToken.Literal[0]] ||
					topToken.Type == token.LparType {
					break
				}
				_, _ = st.Pop()
				postfixTokens = append(postfixTokens, topToken)
			}
			st.Push(t)
		}
	}
	for !st.IsEmpty() {
		// not handled because not empty
		topToken, _ := st.Pop()
		postfixTokens = append(postfixTokens, topToken)
	}
	return postfixTokens, nil
}

func EvaluateRpn(tokens []token.Token) (float64, error) {
	st := stack.New[token.Token]()
	for _, tok := range tokens {
		switch tok.Type {
		case token.NumType:
			st.Push(tok)
		case token.OpType:
			firstTok, ok := st.Pop()
			if !ok {
				return 0, ErrNotEnoughArgs
			}
			secondTok, ok := st.Pop()
			if !ok {
				return 0, ErrNotEnoughArgs
			}
			var firstVal, secondVal float64
			firstVal, err := strconv.ParseFloat(firstTok.Literal, 64)
			if err != nil {
				return 0, fmt.Errorf("failed to parse float %w", err)
			}
			secondVal, err = strconv.ParseFloat(secondTok.Literal, 64)
			if err != nil {
				return 0, fmt.Errorf("failed to parse float %w", err)
			}
			res := token.Actions[tok.Literal[0]](secondVal, firstVal)
			t := token.Token{Type: token.NumType, Literal: strconv.FormatFloat(res, 'f', 3, 64)}
			st.Push(t)
		}
	}
	tok, ok := st.Pop()
	if !ok {
		return 0, ErrNoResult
	}
	result, err := strconv.ParseFloat(tok.Literal, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse float %w", err)
	}
	return result, nil
}
