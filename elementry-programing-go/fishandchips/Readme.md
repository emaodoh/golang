##fishandchips

Description

This is a function called FishAndChips() that takes an int and returns a string.

    If the number is divisible by 2, it will print fish.
    If the number is divisible by 3, it will print chips.
    If the number is divisible by 2 and 3, it will print fish and chips.
    If the number is negative return error: number is negative.
    If the number is non divisible by 2 or 3 return error: non divisible.

Expected function


Usage

Here is a possible program to test the function:

package main

import (
	"fmt"
)

func main() {
	fmt.Println(FishAndChips(4))
	fmt.Println(FishAndChips(9))
	fmt.Println(FishAndChips(6))
}

And its output:

$ go run main.go | cat -e
fish$
chips$
fish and chips$
