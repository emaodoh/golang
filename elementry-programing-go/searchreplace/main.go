package main

import (
	"fmt"
	"os"
)

func SearchReplace(first, second, third string) string {
	var word string
	for _, char := range first {
		if string(char) == second {
			word = word + third
		} else {
			word = word + string(char)
		}
	}
	return word + "\n"
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	first := os.Args[1]
	second := os.Args[2]
	third := os.Args[3]
	if len(second) != 1 || len(third) != 1 {
		return
	}

	word := SearchReplace(first, second, third)
	fmt.Print(word)
}
