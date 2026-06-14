package main

import "fmt"

func Gcd(a, b uint) uint {
	if a == 0 || b == 0 {
		return 0
	}

	var min uint
	var gcd uint

	if a < b {
		min = a
	} else {
		min = b
	}

	for x := uint(1); x <= min; x++ {
		if a%x == 0 && b%x == 0 {
			gcd = x
		}
	}

	return uint(gcd)
}

func main() {
	fmt.Println(Gcd(42, 10))
	fmt.Println(Gcd(42, 12))
	fmt.Println(Gcd(14, 77))
	fmt.Println(Gcd(17, 3))
}
