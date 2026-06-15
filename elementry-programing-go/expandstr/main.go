package main

import (
	"fmt"
	"os"
	"strings"
)

func ExpandStr(text string) string {

	result := strings.Fields(text)
	if len(result) == 0 {
		return ""
	}

	return strings.Join(result, "   ")
}

func main() {

	if len(os.Args) != 2 {
		return
	}
	text := os.Args[1]
	if text == "" {
		return
	}

	result := ExpandStr(text)
	fmt.Println(result)

}
