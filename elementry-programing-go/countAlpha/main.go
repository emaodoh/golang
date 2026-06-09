package main

import (
	"fmt"
)

func CountAlpha(input string) int {
	var count int

	for _, ch := range input {
		if ch >= 'A' && ch <= 'Z' || ch <= 'z' && ch >= 'a' {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountAlpha("Hello world"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1e2l3l4o"))
}
