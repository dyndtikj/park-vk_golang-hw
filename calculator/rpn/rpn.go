package rpn

import (
	"calculator/calculator/token"
	"calculator/stack"
)

// Создает слайс токенов в обратной польской нотации
func CreateRPN(tokens []token.Token) ([]token.Token, bool) {
	postfixTokens := make([]token.Token, 0)
	st := stack.New[token.Token]()
	idxInPostfix := 0
	for _, t := range tokens {
		if t.Type == token.NUMBER {
			postfixTokens = append(postfixTokens, t)
			idxInPostfix++
			continue
		}
		if t.Type == token.L_PAR {
			st.Push(t)
			continue
		}

		if t.Type == token.R_PAR {
			for !st.IsEmpty() {
				topToken, ok := st.Pop()
				if !ok {
					return postfixTokens, ok
				}
				if topToken.Type == token.L_PAR {
					break
				}
				postfixTokens = append(postfixTokens, topToken)
			}
			if st.IsEmpty() {
				// TODO добавить обработку ошибок ( непрвильные скобки )
				return postfixTokens, false
			}
			topToken, ok := st.Top()
			if !ok {
				return postfixTokens, ok
			}
			if topToken.Type == token.L_PAR {
				// ignore err cause handled higher
				_, _ = st.Pop()
			}
			continue
		}
		if t.Type == token.OPERATOR {
			for !st.IsEmpty() {
				topToken, ok := st.Top()
				if !ok {
					return postfixTokens, ok
				}
				if token.Priority[t.Literal] > token.Priority[topToken.Literal] && topToken.Type != token.L_PAR {
					break
				}
				_, _ = st.Pop()
				postfixTokens = append(postfixTokens, topToken)
				idxInPostfix++
			}
			st.Push(t)

			continue
		}
	}
	for !st.IsEmpty() {
		topToken, ok := st.Top()
		if !ok {
			return postfixTokens, ok
		}
		if topToken.Type == token.L_PAR {
			// TODO добавить обработку ошибок ( непрвильные скобки )
			return postfixTokens, false
		} else {
			postfixTokens = append(postfixTokens, topToken)
			_, _ = st.Pop()
			idxInPostfix++
		}
	}
	return postfixTokens, true
}
