package main

import (
	"fmt"
	"strings"
	"unicode"
)

func CamelToSnakeCase(s string) string {
	var output strings.Builder
	if s == "" {
		return s
	}

	for index, char := range s {
		if unicode.IsUpper(char) && index != len(s)-1 {
			if index > 1 && (unicode.IsUpper(rune(s[index-1])) || unicode.IsUpper(rune(s[index+1]))) {
				return s
			}
		}
		if index > 0 && unicode.IsUpper(char) {
			output.WriteString("_")
			output.WriteString(string(char))
		} else {
			output.WriteString(string(char))
		}

		if unicode.IsNumber(char) || unicode.IsPunct(char) {
			return s
		}
	}

	return output.String()
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
