package hw02unpackstring

import (
	"errors"

	"strconv"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	// Place your code here.

	// TODO: wrong cases
	/*
		1. first digit
		2. doubled digits, multiple digits in a row
		3. no letters, but digits
		4. back ticks without digits or escapes
	*/
	// TODO: use cases
	/*
		1. letters only
		2. letters and digits
		3. special characters like \n
		4. empty string
		5. zero (removal of a letter, escaped digit or escape itself)
		6. back ticks and escaped digits or escapes themselves
		7. undocumented: a multi-line string (because of back ticks)
	*/

	if len(input) == 0 {
		return "", nil
	}

	result := []rune{}
	var previousRune rune
	for i, r := range input {
		if isDigit(r) {
			if 0 == i || isDigit(previousRune) {
				return string(result), ErrInvalidString
			}
		}

		if isDigit(r) {
			number, _ := strconv.Atoi(string(r))
			if number > 0 {
				var characters []rune
				for range number - 1 {
					characters = append(characters, previousRune)
				}
				result = append(result, characters...)
			} else {
				result = result[:len(result)-1]
			}
		} else {
			result = append(result, r)
		}

		previousRune = r
	}

	return string(result), nil
}

func isDigit(character rune) bool {
	if _, err := strconv.Atoi(string(character)); err == nil {
		return true
	} else {
		return false
	}
}
