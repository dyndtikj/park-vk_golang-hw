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

func TestOptions_IsValid(t *testing.T) {
	type testCase struct {
		input  Options
		output bool
		name   string
	}

	testCases := []testCase{
		{
			Options{
				CountEntries:   false,
				OnlyRepeating:  false,
				OnlyUnique:     false,
				IgnoreFields:   0,
				IgnoreChars:    0,
				IgnoreRegister: false,
			},
			true,
			"default options",
		},
		{
			Options{
				CountEntries:   true,
				OnlyRepeating:  true,
				OnlyUnique:     false,
				IgnoreFields:   0,
				IgnoreChars:    0,
				IgnoreRegister: false,
			},
			false,
			"multiple -c -d -u not allowed",
		},
		{
			Options{
				CountEntries:   true,
				OnlyRepeating:  true,
				OnlyUnique:     true,
				IgnoreFields:   0,
				IgnoreChars:    0,
				IgnoreRegister: false,
			},
			false,
			"multiple -c -d -u not allowed",
		},
		{
			Options{
				CountEntries:   false,
				OnlyRepeating:  false,
				OnlyUnique:     false,
				IgnoreFields:   -10,
				IgnoreChars:    0,
				IgnoreRegister: false,
			},
			false,
			"negative number of ignored fields",
		},

		{
			Options{
				CountEntries:   false,
				OnlyRepeating:  false,
				OnlyUnique:     false,
				IgnoreFields:   0,
				IgnoreChars:    -10,
				IgnoreRegister: false,
			},
			false,
			"negative number of ignored letters",
		},
		{
			Options{
				CountEntries:   true,
				OnlyRepeating:  false,
				OnlyUnique:     false,
				IgnoreFields:   1,
				IgnoreChars:    10,
				IgnoreRegister: true,
			},
			true,
			"correct options",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.input.IsValid(), test.output)
		})
	}
}

func TestUniq(t *testing.T) {
	type testCase struct {
		input  inputCase
		output []string
		name   string
	}

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
