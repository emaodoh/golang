package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FromTo(from int, to int) string {
	if from < 0 || from > 99 {
		return "Invalid \n"
	}

	if to < 0 || to > 99 {
		return "Invalid \n"
	}
	result := []string{}
	if from > to {

		for x := from; x >= to; x-- {
			num := strconv.Itoa(x)

			result = append(result, "0"+num)
		}
	}
	if from < to {
		for x := from; x <= to; x++ {
			num := strconv.Itoa(x)
			result = append(result, "0"+num)
		}
	}
	if from == to {
		num := strconv.Itoa(to)
		return num + "\n"
	}
	return strings.Join(result, ", ") + "\n"
}

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}
