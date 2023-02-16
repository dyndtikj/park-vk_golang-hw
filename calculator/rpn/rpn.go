package rpn

import (
	"calculator/calculator/token"
	"calculator/stack"
	"errors"
	"strconv"
)

// Создает слайс токенов в обратной польской нотации
func CreateRPN(tokens []token.Token) (postfixTokens []token.Token, err error) {
	st := stack.New[token.Token]()
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
					return postfixTokens, errors.New("internal error, cant create RPN")
				}
				if topToken.Type == token.LparType {
					break
				}
				postfixTokens = append(postfixTokens, topToken)
				_, _ = st.Pop()
			}
			topToken, ok := st.Peek()
			if !ok {
				return postfixTokens, errors.New("internal error, cant create RPN")
			}
			if topToken.Type == token.LparType {
				// ignore err cause handled higher
				_, _ = st.Pop()
			}
		case token.OpType:
			for !st.IsEmpty() {
				topToken, ok := st.Peek()
				if !ok {
					return postfixTokens, errors.New("internal error, cant create RPN")
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
		topToken, ok := st.Pop()
		if !ok {
			return postfixTokens, errors.New("internal error, cant create RPN")
		}
		postfixTokens = append(postfixTokens, topToken)
	}
	return
}

func EvaluateRpn(tokens []token.Token) (result float64, err error) {
	st := stack.New[token.Token]()
	for _, tok := range tokens {
		switch tok.Type {
		case token.NumType:
			st.Push(tok)
		case token.OpType:
			firstTok, ok := st.Pop()
			if !ok {
				return 0, errors.New("Not enough args for operator" + tok.Literal)
			}
			secondTok, ok := st.Pop()
			if !ok {
				return 0, errors.New("Not enough args for operator" + tok.Literal)
			}
			var firstVal, secondVal float64
			firstVal, err = strconv.ParseFloat(firstTok.Literal, 64)
			if err != nil {
				return
			}
			secondVal, err = strconv.ParseFloat(secondTok.Literal, 64)
			if err != nil {
				return
			}
			res := token.Actions[tok.Literal[0]](secondVal, firstVal)
			t := token.Token{Type: token.NumType, Literal: strconv.FormatFloat(res, 'f', 3, 64)}
			st.Push(t)
		}
	}
	tok, ok := st.Pop()
	if !ok {
		return 0, errors.New("expected token of result")
	}
	result, err = strconv.ParseFloat(tok.Literal, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}
