package main

import (
	"fmt"
)

func Atoi(s string) int {
	var result int
	for _, ch := range s {

		if ch >= '0' && ch <= '9' {
			num := int(ch - '0')
			g := result * 10
			result = g + num
		}

	}

	return result
}

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}
