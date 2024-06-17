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
	// for i, r := range input {
	// 	if isDigit(r) {
	// 		if i == 0 || isDigit(previousRune[0]) {
	// 			return string(result), ErrInvalidString
	// 		}
	// 		if isBackslashe(previousRune[0]) && len(previousRune) < 2 {
	// 			result = append(result[:len(result)-1], r)
	// 			previousRune = append(previousRune, r)
	// 			continue
	// 		}
	// 	}
	// 		if i > 0 && isBackslashe(previousRune[0]) && len(previousRune) < 2 {
	// 			// nothing to append to result
	// 			previousRune = append(previousRune, r)
	// 			continue
	// 		}

	// 	if isBackslashe(r) {
	// 		if isBackslashe(previousRune[0]) && len(previousRune) < 2 {
	// 			// nothing to append to result
	// 			previousRune = append(previousRune, r)
	// 			continue
	// 		}
	// 	}

	// 	if isDigit(r) {
	// 		number, _ := strconv.Atoi(string(r))
	// 		if number > 0 {
	// 			var characters []rune
	// 			for range number - 1 {
	// 				if isDigit(previousRune[len(previousRune)-1]) {
	// 				characters = append(characters, previousRune[len(previousRune)-1])
	// 				} else {
	// 					characters = append(characters,previousRune[0])
	// 				}
	// 			}
	// 			result = append(result, characters...)
	// 		} else {
	// 			result = result[:len(result)-1]
	// 		}
	// 		previousRune = []rune{r}
	// 	} else {
	// 		result = append(result, r)
	// 		previousRune = []rune{r}
	// 	}

	// }

	for i, r := range input {
		if isDigit(r) {
			if i == 0 || isDigit(previousRune[0]) {
				return string(result), ErrInvalidString
			}
			if isBackslashe(previousRune[0]) && len(previousRune) < 2 {
				result = append(result[:len(result)-1], r)
				previousRune = append(previousRune, r)
				continue
			}
		}
		if isBackslashe(r) {
			if isBackslashe(previousRune[0]) && len(previousRune) < 2 {
				// nothing to append to result
				previousRune = append(previousRune, r)
				continue
			}
		}
		if i > 0 && isBackslashe(previousRune[0]) && len(previousRune) < 2 {
			// nothing to append to result
			previousRune = append(previousRune, r)
			continue
		}

		if isDigit(r) {
			number, _ := strconv.Atoi(string(r))
			if number > 0 {
				var characters []rune
				for range number - 1 {
					if isDigit(previousRune[len(previousRune)-1]) {
						characters = append(characters, previousRune[len(previousRune)-1])
					} else {
						characters = append(characters, previousRune[0])
					}
				}
				result = append(result, characters...)
			} else {
				result = result[:len(result)-1]
			}
			previousRune = []rune{r}
		} else {
			result = append(result, r)
			previousRune = []rune{r}
		}
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
