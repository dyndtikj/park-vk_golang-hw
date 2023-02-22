package parser

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"homework/hw1_part2/calculator/token"
	"testing"
)

func TestParse(t *testing.T) {
	type outputCase struct {
		tokens []token.Token
		err    error
	}
	type testCase struct {
		input  string
		output outputCase
		name   string
	}

	testCases := []testCase{
		{
			"1+2",
			outputCase{
				tokens: []token.Token{{token.NumType, "1"},
					{token.OpType, "+"},
					{token.NumType, "2"}},
				err: nil,
			},
			"test : 1+2 ",
		},
		{
			"111/2121*1-10",
			outputCase{
				tokens: []token.Token{{token.NumType, "111"},
					{token.OpType, "/"},
					{token.NumType, "2121"},
					{token.OpType, "*"},
					{token.NumType, "1"},
					{token.OpType, "-"},
					{token.NumType, "10"}},
				err: nil,
			},
			"test : 111/2121*1-10",
		},
		{
			"(1+2)*12",
			outputCase{
				tokens: []token.Token{{token.LparType, "("},
					{token.NumType, "1"},
					{token.OpType, "+"},
					{token.NumType, "2"},
					{token.RparType, ")"},
					{token.OpType, "*"},
					{token.NumType, "12"}},
				err: nil,
			},
			"test : (1+2)*12",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			tokens, err := Parse(test.input)
			if err != nil && test.output.err != nil {
				if err.Error() != test.output.err.Error() {
					t.Errorf("Expected other err, then %v ", err)
				}
			} else if err != test.output.err {
				fmt.Println(err.Error(), test.output.err.Error())
				t.Errorf("Expected other err, then %v ", err)
			}
			for i, tok := range tokens {
				assert.Equal(t, test.output.tokens[i].Literal, tok.Literal)
				assert.Equal(t, test.output.tokens[i].Type, tok.Type)
			}
		})
	}
}
