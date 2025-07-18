package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("enter a number: ")
	fmt.Scan(&num)
	if num < 1 || num <= 100 {
		fmt.Print("norm")
	} else {
		fmt.Print("many")
	}
}
