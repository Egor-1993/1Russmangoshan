package main

import "fmt"

func main() {
	var a int
	fmt.Printf("enter a: ")
	fmt.Scan(&a)
	lastDigit := a % 10
	if lastDigit < 0 {
		lastDigit = -lastDigit
	}
	fmt.Printf("%d\n", lastDigit)
}
