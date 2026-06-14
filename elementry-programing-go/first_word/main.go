package main

import (
	"strings"
)

func FirstWord(s string) string {
	if len(s) == 0 {
		return "" + "\n"
	}
	sen := strings.Fields(s)
	return sen[0] + "\n"
}
