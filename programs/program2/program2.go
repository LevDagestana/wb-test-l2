package main

import (
	"errors"
	"fmt"
)

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func isSlash(c rune) bool {
	return c == '\\'
}

func checker(s string) error {

	if isDigit(rune(s[0])) {
		return errors.New("некорректная строка")
	}
	return nil
}

func Unpack(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	if err := checker(s); err != nil {
		return "", err
	}

	var result []rune
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if isSlash(runes[i]) {

			result = append(result, runes[i+1])

		} else if isDigit(runes[i]) {
			if len(result) == 0 {
				return "", errors.New("некорректная строка")
			}

			count := int(runes[i] - '0')
			if count > 0 {
				lastChar := result[len(result)-1]
				for j := 1; j < count; j++ {
					result = append(result, lastChar)
				}
			}
		} else {
			result = append(result, runes[i])
		}
	}

	return string(result), nil
}

func main() {

	tests := []string{
		"a4bc2d5e",
		"abcd",
		"45",
		"",
		`qwe\4\5`,
		`qwe\45`,
		`qwe\\5`,
	}

	for _, test := range tests {
		result, err := Unpack(test)
		if err != nil {
			fmt.Printf("Ошибка '%s': %v\n", test, err)
		} else {
			fmt.Printf("Распаковка '%s': %s\n", test, result)
		}
	}
}
