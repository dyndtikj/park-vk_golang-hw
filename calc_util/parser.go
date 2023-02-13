package calc_util

func Parse(input string) (result []Token, err error) {
	tokenizer := NewTokenizer(input)

	for {
		var token Token
		token, err = tokenizer.NextToken()
		if err != nil {
			return
		}
		if token.Type == EOF {
			break
		}
		result = append(result, token)
	}
	return
}
