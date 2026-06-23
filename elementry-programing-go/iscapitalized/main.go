package main

import (
	"strings"
	"unicode"
)

func IsCapitalized(s string) bool {
	if s == "" {
		return false
	}
	var result bool
	sen := strings.Fields(s)
	for _, word := range sen {

		if !unicode.IsLetter(rune(word[0])) || unicode.IsUpper(rune(word[0])) {
			result = true
		} else {
			return false
		}

	}
	return result
}
