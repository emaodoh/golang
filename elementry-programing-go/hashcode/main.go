package main

import "fmt"

func HashCode(dec string) string {
	word := ""
	for _, char := range dec {
		hash := (char + rune(len(dec))) % 127

		if hash < 33 {
			hash += 33
		}
		word = word + string(hash)
	}
	return word
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
