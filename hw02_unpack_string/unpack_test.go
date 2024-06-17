package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		// more cases
		{input: "abc1ABC", expected: "abcABC"},
		// uncomment if task with asterisk completed
		{input: `qwe\4\5`, expected: `qwe45`},
		{input: `qwe\45`, expected: `qwe44444`},
		{input: `qwe\\5`, expected: `qwe\\\\\`},
		{input: `qwe\\\3`, expected: `qwe\3`},
		// from README.ms
		{input: `d\n5abc`, expected: `d\n\n\n\n\nabc`},
		// more special characters
		// https://en.wikipedia.org/wiki/List_of_Unicode_characters
		{input: `d\z2abc`, expected: `d\z\zabc`},
		{input: `d\b3abc`, expected: `d\b\b\babc`},
		{input: `d\a0abc`, expected: `dabc`},
		//
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

func TestIsDigit(t *testing.T) {
	testData := []struct {
		input       []rune
		expected    bool
		description string
	}{
		{input: []rune("3"), expected: true, description: "digit"},
		{input: []rune(`\0`), expected: false, description: "empty"},
		{input: []rune("a"), expected: false, description: "letter"},
		{input: []rune("\\"), expected: false, description: "backslashe"},
	}
	for _, tc := range testData {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			actual := isDigit(tc.input[0])
			require.Equal(t, tc.expected, actual)
		})
	}
}

func TestIsBackslashe(t *testing.T) {
	testData := []struct {
		input       []rune
		expected    bool
		description string
	}{
		{input: []rune("3"), expected: false, description: "digit"},
		{input: []rune(`\`), expected: true, description: "backshashe in back ticks"},
		{input: []rune("a"), expected: false, description: "letter"},
		{input: []rune("\\"), expected: true, description: "backslashe in quotes"},
	}
	for _, tc := range testData {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			actual := isBackslashe(tc.input[0])
			require.Equal(t, tc.expected, actual)
		})
	}
}

func TestIsLetter(t *testing.T) {
	testData := []struct {
		input       []rune
		expected    bool
		description string
	}{
		{input: []rune("3"), expected: false, description: "digit"},
		{input: []rune(`\`), expected: false, description: "backshashe in back ticks"},
		{input: []rune("a"), expected: true, description: "letter"},
		{input: []rune("\\"), expected: false, description: "backslashe in quotes"},
	}
	for _, tc := range testData {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			actual := isLetter(tc.input[0])
			require.Equal(t, tc.expected, actual)
		})
	}
}

func TestFirstInSlkice(t *testing.T) {
	testData := []struct {
		input       []rune
		expected    rune
		description string
	}{
		{input: []rune{rune(70), rune(71), rune(72)}, expected: rune(70), description: "multiple elements"},
		{input: []rune{rune(70)}, expected: rune(70), description: "single element"},
	}
	for _, tc := range testData {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			actual := first(tc.input)
			require.Equal(t, tc.expected, actual)
		})
	}
}

func TestLastInSlkice(t *testing.T) {
	testData := []struct {
		input       []rune
		expected    rune
		description string
	}{
		{input: []rune{rune(70), rune(71), rune(72)}, expected: rune(72), description: "multiple elements"},
		{input: []rune{rune(70)}, expected: rune(70), description: "single element"},
	}
	for _, tc := range testData {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			actual := last(tc.input)
			require.Equal(t, tc.expected, actual)
		})
	}
}
