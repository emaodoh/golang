This is a function called RetainFirstHalf() that takes a string as an argument and returns the first half of this string.

    If the length of the string is odd, it will it down.
    If the string is empty, it will return an empty string.
    If the string length is equal to one, it will return the string.


Usage

Here is a possible program to test the function:

package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(piscine.RetainFirstHalf("A"))
	fmt.Println(piscine.RetainFirstHalf(""))
	fmt.Println(piscine.RetainFirstHalf("Hello World"))
}

And its output:

$ go run main.go | cat -e
This is the 1st half$
A$
$
Hello$
