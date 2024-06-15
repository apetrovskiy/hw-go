package hw02unpackstring

import (
	"errors"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(_ string) (string, error) {
	// Place your code here.

	// TODO: wrong cases
	/*
		1. first digit
		2. doubled digits, multiple digits
		3. no letters
		4. back ticks without digits or escapes
	*/
	// TODO: use cases
	/*
		1. letters only
		2. letters and digits
		3. special characters like \n
		4. empty string
		5. zero (removal of a letter)
		6. back ticks and escaped digits or escapes
	*/
	return "", nil
}
