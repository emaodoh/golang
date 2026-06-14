package main

import (
	"fmt"
	"strings"
)

func RepeatAlpha(s string) string {
	var word string
	for _, char := range s {
		if char >= 'a' && char <= 'z' {

			word = word + strings.Repeat(string(char), int(char-'a')+1)
		} else if char >= 'A' && char <= 'Z' {
			word = word + strings.Repeat(string(char), int(char-'A')+1)

		} else {
			word = word + string(char)
		}
	}
	return word
}

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}
