package main

import (
	"fmt"
	"strings"
)

func LastWord(s string) string {

	sen := strings.Fields(s)

	if len(sen) == 0 {
		return "\n"
	}

	return sen[len(sen)-1] + "\n"
}

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}
