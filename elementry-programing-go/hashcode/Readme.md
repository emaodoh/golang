hashcode
Description

This is a function called HashCode() that takes a string as an argument and returns a new hashed string.

    The hash equation is computed as follows:

(ASCII of current character + size of the string) % 127, ensuring the result falls within the ASCII range of 0 to 127.

    If the resulting character is unprintable i will add 33 to it.


Usage

Here is a possible program to test the function:

package main

import (
	"fmt"
)   

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}

And its output:

$ go run main.go
B
CD
EDF
Spwwz+bz}wo
