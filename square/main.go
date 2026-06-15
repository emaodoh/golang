package main

import "fmt"

func main() {
	var numColu int
	var numrow int
	var symbol string

	fmt.Print("Enter number of row: ")
	fmt.Scan(&numrow)

	fmt.Print("Enter number of COLUMN: ")
	fmt.Scan(&numColu)

	fmt.Print("Enter number of SYMBOL: ")
	fmt.Scan(&symbol)

	for i := 1; i <= numColu; i++ { 
		for j := 1; j <= numrow; j++ {
			fmt.Printf("%v ", symbol)
		}
		fmt.Println()
	}
}
