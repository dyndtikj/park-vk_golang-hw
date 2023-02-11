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
			"-c (first test)",
		},
		{
			inputCase{[]string{
				"Hello", "Bye", "Hi", "Hi", "\n", "Bonjour", "Bye", "Thanks", "Hi"},
				Options{CountEntries: true}},
			[]string{
				"1 Hello", "1 Bye", "2 Hi", "1 \n", "1 Bonjour", "1 Bye", "1 Thanks", "1 Hi"},
			"-c (second test)",
		},
		{
			inputCase{[]string{
				"Hello", "Bye", "Hi", "Hi", "\n", "Bonjour", "Bye", "Thanks", "Hi"},
				Options{OnlyUnique: true}},
			[]string{
				"Hello", "Bye", "\n", "Bonjour", "Bye", "Thanks", "Hi"},
			"-u",
		},
		{
			inputCase{[]string{
				"Hello", "Bye", "Hi", "Hi", "\n", "Bonjour", "Bye", "Thanks", "Hi"},
				Options{OnlyRepeating: true}},
			[]string{"Hi"},
			"-d",
		},
		{
			inputCase{[]string{
				"I LOVE MUSIC.", "I love music.", "I love music.", "\n",
				"I love MuSIC of Kartik.", "I love music of Kartik.", "Thanks.",
				"I love music of Kartik.", "I love MuSIC of Kartik."},
				Options{IgnoreRegister: true}},
			[]string{
				"I LOVE MUSIC.", "\n", "I love MuSIC of Kartik.",
				"Thanks.", "I love music of Kartik."},
			"-i",
		},
		{
			inputCase{[]string{
				"We love music.", "I love music.", "They love music.", "\n",
				"I love music of Kartik.", "We love music of Kartik.",
				"Thanks."},
				Options{IgnoreFields: 1}},
			[]string{
				"We love music.", "\n", "I love music of Kartik.",
				"Thanks."},
			"-f 1",
		},
		{
			inputCase{[]string{
				"I love music.", "A love music.", "C love music.", "\n",
				"I love music of Kartik.", "We love music of Kartik.", "Thanks."},
				Options{IgnoreChars: 1}},
			[]string{
				"I love music.", "\n", "I love music of Kartik.",
				"We love music of Kartik.", "Thanks."},
			"-s 1",
		},
		{
			inputCase{[]string{
				"We love music.", "I love music.", "They hate music.",
				"I love music of Kartik.", "We have music of Kartik.",
				"They have music of Kartik."},
				Options{IgnoreFields: 1, IgnoreChars: 3}},
			[]string{
				"We love music.", "They hate music.",
				"I love music of Kartik.", "We have music of Kartik."},
			"-f 1 -s 3",
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
