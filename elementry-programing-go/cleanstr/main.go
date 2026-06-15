package main

import (
	"fmt"
	"os"
	"strings"
)

func CleanStr(text string) string {
	if len(text) == 0 {
		return ""
	}

	result := strings.Fields(text)

	return strings.Join(result, " ")
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}
	text := os.Args[1]

	result := CleanStr(text)

	fmt.Println(result)
}
