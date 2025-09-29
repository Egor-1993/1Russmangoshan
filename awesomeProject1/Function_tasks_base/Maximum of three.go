package main

import "fmt"

func max3(a, b, c int) int {
	if a > b && a > c {
		return a
	} else if b > a && b > c {
		return b
	} else {
		return c
	}

}
func main() {
	var a, b, c int
	fmt.Print("Enter first number: ")
	fmt.Scan(&a)
	fmt.Print("Enter second number: ")
	fmt.Scan(&b)
	fmt.Print("Enter third number: ")
	fmt.Scan(&c)
	result := max3(a, b, c)
	fmt.Println("maximum of three", result)

}
