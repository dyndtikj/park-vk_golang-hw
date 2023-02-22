package rpn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"homework/hw1_part2/calculator/token"
	"testing"
)

func TestEvaluateRpn(t *testing.T) {
	type outputCase struct {
		result float64
		err    error
	}
	type testCase struct {
		input  []token.Token
		output outputCase
		name   string
	}

	testCases := []testCase{
		{
			[]token.Token{
				{token.NumType, "1"},
				{token.NumType, "2"},
				{token.OpType, "+"}},
			outputCase{
				result: 3,
				err:    nil,
			},
			"test : 1+2 = 3",
		},
		{
			// 2+3*4
			[]token.Token{
				{token.NumType, "2"},
				{token.NumType, "3"},
				{token.NumType, "4"},
				{token.OpType, "*"},
				{token.OpType, "+"}},
			// 234*+
			outputCase{
				result: 14,
				err:    nil,
			},
			"test : 2+3*4 = 14",
		},
		{
			[]token.Token{
				{token.NumType, "2"},
				{token.NumType, "3"},
				{token.OpType, "+"},
				{token.NumType, "4"},
				{token.OpType, "*"}},
			outputCase{
				result: 20,
				err:    nil,
			},
			"test : (2+3)*4 = 20",
		},
		{
			[]token.Token{
				{token.NumType, "111"},
				{token.NumType, "121"},
				{token.OpType, "+"},
				{token.NumType, "91"},
				{token.NumType, "23"},
				{token.OpType, "-"},
				{token.OpType, "*"}},
			outputCase{
				result: 15776,
				err:    nil,
			},
			"test :(111+121)*(91-23) = 15776",
		},
		{
			// 111+121*91-23
			[]token.Token{
				{token.NumType, "111"},
				{token.NumType, "121"},
				{token.NumType, "91"},
				{token.OpType, "*"},
				{token.OpType, "+"},
				{token.NumType, "23"},
				{token.OpType, "-"}},

			outputCase{
				result: 11099,
				err:    nil,
			},
			"test : 111+121*91-23 = 11099",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			res, err := EvaluateRpn(test.input)
			if err != nil && test.output.err != nil {
				if err.Error() != test.output.err.Error() {
					t.Errorf("Expected other err, then %v ", err)
				}
			} else if err != test.output.err {
				fmt.Println(err.Error(), test.output.err.Error())
				t.Errorf("Expected other err, then %v ", err)
			}
			assert.Equal(t, test.output.result, res)
		})
	}
}

func TestCreateRPN(t *testing.T) {
	type outputCase struct {
		tokens []token.Token
		err    error
	}
	type testCase struct {
		input  []token.Token
		output outputCase
		name   string
	}

	testCases := []testCase{
		{
			// 1+2
			[]token.Token{
				{token.NumType, "1"},
				{token.OpType, "+"},
				{token.NumType, "2"}},
			// 12+
			outputCase{
				tokens: []token.Token{
					{token.NumType, "1"},
					{token.NumType, "2"},
					{token.OpType, "+"}},
				err: nil,
			},
			"test : 1+2 to 12+",
		},
		{
			// 2+3*4
			[]token.Token{
				{token.NumType, "2"},
				{token.OpType, "+"},
				{token.NumType, "3"},
				{token.OpType, "*"},
				{token.NumType, "4"}},
			// 234*+
			outputCase{
				tokens: []token.Token{
					{token.NumType, "2"},
					{token.NumType, "3"},
					{token.NumType, "4"},
					{token.OpType, "*"},
					{token.OpType, "+"}},
				err: nil,
			},
			"test : 2+3*4 to (2 3 4*+)",
		},
		{
			// (2+3)*4
			[]token.Token{
				{token.LparType, "("},
				{token.NumType, "2"},
				{token.OpType, "+"},
				{token.NumType, "3"},
				{token.RparType, ")"},
				{token.OpType, "*"},
				{token.NumType, "4"}},
			// 23+4*
			outputCase{
				tokens: []token.Token{
					{token.NumType, "2"},
					{token.NumType, "3"},
					{token.OpType, "+"},
					{token.NumType, "4"},
					{token.OpType, "*"}},
				err: nil,
			},
			"test : (2+3)*4 to (2 3 + 4 *)",
		},
		{
			// (111+121)*(91-23)
			[]token.Token{
				{token.LparType, "("},
				{token.NumType, "111"},
				{token.OpType, "+"},
				{token.NumType, "121"},
				{token.RparType, ")"},
				{token.OpType, "*"},
				{token.LparType, "("},
				{token.NumType, "91"},
				{token.OpType, "-"},
				{token.NumType, "23"},
				{token.RparType, ")"}},
			// 111 121 + 91 23 - *
			outputCase{
				tokens: []token.Token{
					{token.NumType, "111"},
					{token.NumType, "121"},
					{token.OpType, "+"},
					{token.NumType, "91"},
					{token.NumType, "23"},
					{token.OpType, "-"},
					{token.OpType, "*"}},
				err: nil,
			},
			"test :(111+121)*(91-23) to (111 121 + 91 23 - *)",
		},
		{
			// 111+121*91-23
			[]token.Token{
				{token.NumType, "111"},
				{token.OpType, "+"},
				{token.NumType, "121"},
				{token.OpType, "*"},
				{token.NumType, "91"},
				{token.OpType, "-"},
				{token.NumType, "23"}},
			// 111 121 91 * + 23 -
			outputCase{
				tokens: []token.Token{
					{token.NumType, "111"},
					{token.NumType, "121"},
					{token.NumType, "91"},
					{token.OpType, "*"},
					{token.OpType, "+"},
					{token.NumType, "23"},
					{token.OpType, "-"}},
				err: nil,
			},
			"test :111+121*91-23 to (111 121 91 * + 23 -)",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			tokens, err := CreateRPN(test.input)
			if err != test.output.err {
				t.Errorf("Expected result of operation: %v, got %v ", test.output.err, err)
			}
			for i, tok := range tokens {
				assert.Equal(t, test.output.tokens[i].Literal, tok.Literal)
				assert.Equal(t, test.output.tokens[i].Type, tok.Type)
			}
		})
	}
}
