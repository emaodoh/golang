
This is a function that takes a string and a character as arguments and returns the number of times the character appears in the string.

    if the character is not in the string it returns 0
    if the string is empty it returns 0

Usage

Here is a possible program to test the function:

package main

import (
	"fmt"
)

func main() {
	fmt.Println(CountChar("Hello World", 'l'))
	fmt.Println(CountChar("5  balloons", 5))
	fmt.Println(CountChar("   ", ' '))
	fmt.Println(CountChar("The 7 deadly sins", '7'))
}

And its output :

$ go run .
3
0
3
1
