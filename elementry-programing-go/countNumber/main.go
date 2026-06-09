package main

import (
	"fmt"
)

func CheckNumber(input string) bool {
	for _, ch := range input {
		if ch >= '1' && ch <= '9' {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
}
