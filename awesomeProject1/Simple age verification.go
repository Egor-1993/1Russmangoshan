package main

import (
	"fmt"
)

func main() {
	var age int
	fmt.Print("How old are you")
	fmt.Scan(&age)
	if age <= 7 {
		fmt.Println("you're very small")

	} else {
		fmt.Println("you're very big")
	}
}
