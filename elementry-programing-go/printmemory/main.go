package main

import "fmt"

// func PrintMemory(arr [10]byte) {
// 	st := ""
// 	count := 0
// 	for _, char := range arr {
// 		ch := rune(char)
// 		if ch < ' ' || ch > '~' {
// 			st = st + "."
// 		} else {
// 			st = st + string(ch)

// 		}
// 		fmt.Printf("%x ", ch)
// 		count++
// 		if count == 4 {
// 			fmt.Println()
// 			count = 0
// 		}

// 	}

// 	fmt.Println(st)

// }

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func PrintMemory(arr [10]byte) {
	st := ""
	count := 0

	for _, char := range arr {
		ch := rune(char)

		if ch < ' ' || ch > '~' {
			st += "."
		} else {
			st += string(ch)
		}

		fmt.Printf("%02x ", ch)
		count++

		if count == 4 {
			fmt.Println()
			count = 0
		}
	}

	// If we're still on the same line, move to the next line.
	if count != 0 {
		fmt.Println()
	}

	fmt.Println(st)
}
