This is a function that converts a string from camelCase to snake_case.

    If the string is empty, it will return an empty string.
    If the string is not camelCase, it will return the string unchanged.
    If the string is camelCase, it will return the snake_case version of the string.

Note that camelCase has two different writing alternatives that will be accepted:

    lowerCamelCase
    UpperCamelCase

Rules for writing in camelCase:

    The word does not end on a capitalized letter (CamelCasE).
    No two capitalized letters shall follow directly each other (CamelCAse).
    Numbers or punctuation are not allowed in the word anywhere (camelCase1).

Usage

Here is a possible program to test your function:

package main

import (
	"fmt"
)

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}

And its output:

$ go run .
Hello_World
hello_World
camel_Case
CAMELtoSnackCASE
camel_To_Snake_Case
hey2
