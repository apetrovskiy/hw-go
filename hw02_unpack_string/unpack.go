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
	var previousRune []rune

	for i, r := range input {
		if isDigit(r) {
			// 1xxx; xxx11xxx -> error
			if i == 0 || isDigit(first(previousRune)) {
				return string(result), ErrInvalidString
			}
			// \1 -> result + 1, previous \1
			if isBackslashe(first(previousRune)) && len(previousRune) == 1 {
				result = append(result[:len(result)-1], r)
				previousRune = append(previousRune, r)
				continue
			}
		}
		// \a -> privous \a
		if i > 0 && isBackslashe(first(previousRune)) && len(previousRune) == 1 {
			// nothing to append to result
			previousRune = append(previousRune, r)
			continue
		}

		if isDigit(r) {
			number, _ := strconv.Atoi(string(r))
			if number == 0 {
				result = result[:len(result)-1]
				continue
			}
			if isLetter(last(previousRune)) && isBackslashe(first(previousRune)) {
				number++
				result = result[:len(result)-1]
			}
			var characters []rune
			for range number - 1 {
				if isLetter(last(previousRune)) {
					characters = append(characters, previousRune...)
				} else {
					characters = append(characters, last(previousRune))
				}
			}
			result = append(result, characters...)
			previousRune = []rune{r}
			continue
		}
		result = append(result, r)
		previousRune = []rune{r}
	}

	return string(result), nil
}

func isDigit(character rune) bool {
	if string(character) == "" {
		return false
	}
	if _, err := strconv.Atoi(string(character)); err == nil {
		return true
	}
	return false
}

func isBackslashe(character rune) bool {
	return string(character) == `\`
}

func isLetter(character rune) bool {
	return !isDigit(character) && !isBackslashe(character)
}

func first(characters []rune) rune {
	return characters[0]
}

func last(characters []rune) rune {
	return characters[len(characters)-1]
}
