package calculator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculate(t *testing.T) {
	type outputCase struct {
		result float64
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
				result: 3,
				err:    nil,
			},
			"test : 1+2 = 3",
		},
		{
			"2+2*2",
			outputCase{
				result: 6,
				err:    nil,
			},
			"test : 2+2*2 = 6",
		},
		{
			"(2+2)*2",
			outputCase{
				result: 8,
				err:    nil,
			},
			"test : (2+2)*2 = 8",
		},
		{
			"(8+2)/4",
			outputCase{
				result: 2.5,
				err:    nil,
			},
			"test : (8+2)/4 = 2.5",
		},
		{
			"0/2*100",
			outputCase{
				result: 0,
				err:    nil,
			},
			"test : 0/2*100 = 0",
		},
		{
			"1",
			outputCase{
				result: 1,
				err:    nil,
			},
			"test : 1 = 1",
		},
		{
			"8+2/2",
			outputCase{
				result: 9,
				err:    nil,
			},
			"test : 8+2/2 = 9",
		},
		{
			"1/3",
			outputCase{
				result: 0.333,
				err:    nil,
			},
			"test : 1/3 = 0.333",
		},
		{
			"(111+333)/111+10*(99+2)",
			outputCase{
				result: 1014,
				err:    nil,
			},
			"test : (111+333)/111+10*(99+2) = 1014",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			res, err := Calculate(test.input)
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
