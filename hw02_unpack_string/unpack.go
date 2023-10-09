package hw02unpackstring

import (
	"errors"
	"strconv"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (result string, err error) {

	var slash bool

	for i, symbol := range str {

		if symbol == '\\' && !slash {
			slash = true
			continue
		}

		if slash {
			if unicode.IsLetter(symbol) {
				return "", ErrInvalidString
			}

			result += string(symbol)
			slash = false
			continue
		}

		if unicode.IsDigit(symbol) {
			if i == 0 {
				return "", ErrInvalidString
			}

			if unicode.IsDigit([]rune(str)[i-1]) && []rune(str)[i-2] != '\\' {
				return "", ErrInvalidString
			}

			val, _ := strconv.Atoi(string(symbol))
			if val == 0 {
				result = result[:len(result)-1]
				continue
			}

			for j := 0; j < val-1; j++ {
				result += string([]rune(str)[i-1])
			}
			continue
		}

		result += string(symbol)
	}

	return result, nil
}
