package rpn

import (
	"calculator/calculator/token"
	"calculator/stack"
	"errors"
	"strconv"
)

// Создает слайс токенов в обратной польской нотации
func CreateRPN(tokens []token.Token) ([]token.Token, bool) {
	postfixTokens := make([]token.Token, 0)
	st := stack.New[token.Token]()
	for _, t := range tokens {
		switch t.Type {
		case token.NUMBER:
			postfixTokens = append(postfixTokens, t)
		case token.L_PAR:
			st.Push(t)
		case token.R_PAR:
			for !st.IsEmpty() {
				topToken, ok := st.Top()
				if !ok {
					return postfixTokens, ok
				}
				if topToken.Type == token.L_PAR {
					break
				}
				postfixTokens = append(postfixTokens, topToken)
				_, _ = st.Pop()
			}
			topToken, ok := st.Top()
			if !ok {
				return postfixTokens, ok
			}
			if topToken.Type == token.L_PAR {
				// ignore err cause handled higher
				_, _ = st.Pop()
			}
		case token.OPERATOR:
			for !st.IsEmpty() {
				topToken, ok := st.Top()
				if !ok {
					return postfixTokens, ok
				}
				if token.Priority[t.Literal] > token.Priority[topToken.Literal] || topToken.Type == token.L_PAR {
					break
				}
				_, _ = st.Pop()
				postfixTokens = append(postfixTokens, topToken)
			}
			st.Push(t)
		}
	}
	for !st.IsEmpty() {
		topToken, ok := st.Top()
		if !ok {
			return postfixTokens, ok
		}
		postfixTokens = append(postfixTokens, topToken)
		_, _ = st.Pop()
	}
	return postfixTokens, true
}

func EvaluateRpn(tokens []token.Token) (float64, error) {
	st := stack.New[token.Token]()
	for _, tok := range tokens {
		switch tok.Type {
		case token.NUMBER:
			st.Push(tok)
		case token.OPERATOR:
			firstTok, ok := st.Pop()
			if !ok {
				return 0, errors.New("Not enough args for operator" + tok.Literal)
			}
			secondTok, ok := st.Pop()
			if !ok {
				return 0, errors.New("Not enough args for operator" + tok.Literal)
			}
			// TODO fix this :-)
			var res float64
			firstVal, err := strconv.ParseFloat(firstTok.Literal, 64)
			if err != nil {
				return 0, errors.New("Cant convert token: " + firstTok.Literal + "to number")
			}
			secondVal, err := strconv.ParseFloat(secondTok.Literal, 64)
			if err != nil {
				return 0, errors.New("Cant convert token: " + secondTok.Literal + "to number")
			}
			switch tok.Literal {
			case "+":
				res = secondVal + firstVal
			case "-":
				res = secondVal - firstVal
			case "*":
				res = secondVal * firstVal
			case "/":
				res = secondVal / firstVal
			}
			t := token.NewToken(token.NUMBER, strconv.FormatFloat(res, 'f', 3, 64))
			st.Push(t)
		}
	}
	tok, ok := st.Pop()
	if !ok {
		return 0, errors.New("Expected token of result")
	}
	result, err := strconv.ParseFloat(tok.Literal, 64)
	if err != nil {
		return 0, errors.New("Cant convert token: " + tok.Literal + "to number")
	}
	return result, nil
}
