package main

import (
	"fmt"
)

// func Atoi(s string) int {
// 	var result int
// 	for index, ch := range s {
// 		if ch == ' ' {
// 			return 0
// 		}

// 		if ch >= '0' && ch <= '9' {
// 			num := int(ch - '0')
// 			g := result * 10
// 			result = g + num
// 		}

// 		if index == 0 && ch == '-' {
// 			result = -result

// 		}

// 	}
// 	return result
// }

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	result := 0
	sign := 1
	startIndex := 0
	switch s[0] {
	case '-':
		sign = -1
		startIndex = 1
	case '+':
		startIndex = 1

	}

	for i := startIndex; i < len(s); i++ {
		ch := s[i]

		if ch < '0' || ch > '9' {
			return 0
		}

		num := int(ch - '0')
		result = result*10 + num
	}

	return result * sign
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
