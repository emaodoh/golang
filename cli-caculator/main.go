package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 4 {
		fmt.Println("unsupported input")
		return
	}

	nums1 := os.Args[1]
	nums2 := os.Args[3]
	operator := os.Args[2]

	num1, err := strconv.Atoi(nums1)
	if err != nil {
		fmt.Println("numbers only")
		return
	}
	num2, err := strconv.Atoi(nums2)
	if err != nil {
		fmt.Println("numbers only")
		return
	}
	switch operator {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		fmt.Println(num1 / num2)
	default:
		fmt.Println("invalid operator")
	}

}
