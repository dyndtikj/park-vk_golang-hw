package uniq

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type inputCase struct {
	lines   []string
	options Options
}

type testCase struct {
	input  inputCase
	output []string
	name   string
}

func TestUniq(t *testing.T) {
	testCases := []testCase{
		{
			inputCase{[]string{"I love Go", "I love Go", "You love Go", "You love Go", "You love Go", "You love Go", "Rob Pike loves Go"},
				Options{CountEntries: true}},
			[]string{"2 I love Go", "4 You love Go", "1 Rob Pike loves Go"},
			"test1",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result, err := Uniq(test.input.options, test.input.lines)
			if err != nil {
				log.Fatal(err)
			}
			assert.Equal(t, result, test.output)
		})
	}
}
