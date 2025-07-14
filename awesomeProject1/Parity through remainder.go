package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Enter number: ")
	fmt.Scan(&num)
	fmt.Print(num % 2)
}
