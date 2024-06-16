package hw02unpackstring

import (
	"errors"
	"fmt"
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

	// onlyDigits := true
	result := []rune{}
	var previousRune rune
	// var builder = strconv.Builder
	for i, r := range input {
		fmt.Printf("========= index = %d, rune = %s \n", i, string(r))
		if isDigit(r) && isDigit(previousRune) {
			return "", ErrInvalidString
		}

		if isDigit(r) {
			number, _ := strconv.Atoi(string(r))
			var characters []rune
			for range number {
				characters = append(characters, previousRune)
			}
			result = append(result, characters...)
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
