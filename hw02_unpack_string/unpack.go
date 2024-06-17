package hw02unpackstring

import (
	"errors"
	"strconv"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	if len(input) == 0 {
		return "", nil
	}

	result := []rune{}
	previousRune := []rune{}

	for i, r := range input {
		// \a -> previous \a
		if len(previousRune) == 1 && isBackslashe(first(previousRune)) {
			if isDigit(r) {
				result = append(result[:len(result)-1], r)
				previousRune = append(previousRune, r)
				continue
			}
			// nothing to append to result
			previousRune = append(previousRune, r)
			continue
		}
		if isDigit(r) && (i == 0 || isDigit(first(previousRune))) {
			return string(result), ErrInvalidString
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
