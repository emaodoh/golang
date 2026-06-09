package main

func CountChar(input string, w rune) int {
	var count int
	for _, word := range input {
		if word == w {
			count++
		}
	}
	return count
}
