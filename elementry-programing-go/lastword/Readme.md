lastword

Description

This is a function LastWord that takes a string and returns its last word followed by a \n.

    A word is a section of string delimited by spaces or by the start/end of the string.

Usage

Here is a possible program to test the function :

package main

import (
	"fmt"
	
)

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}

And its output :

$ go run main.go | cat -e
not$
lorem,ipsum$
$
$

