package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Print("enter a")
	fmt.Scan(&a)
	fmt.Print("enter b")
	fmt.Scan(&b)
	if a > b {
		fmt.Print("a >b")
	} else if b > a {
		fmt.Print("b >a")

	} else {
		fmt.Print("a =b")
	}
}
