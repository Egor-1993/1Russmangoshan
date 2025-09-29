package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Print("Enter a: ")
	fmt.Scan(&a)
	fmt.Print("Enter b: ")
	fmt.Scan(&b)
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	fmt.Print(diff)

}
