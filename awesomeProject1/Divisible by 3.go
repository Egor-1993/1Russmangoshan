package main

import (
	"fmt"
)

func main() {
	var numb int
	fmt.Print("Enter your number: ")
	fmt.Scan(&numb)

	if numb%3 == 0 {
		fmt.Println("Divisible by 3")
	} else {
		fmt.Println("Not divisible by 3")
	}
}
