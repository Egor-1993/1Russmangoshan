package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Enter a number: ")
	fmt.Scan(&a)
	fmt.Print("Enter power of a number a: ")
	fmt.Scan(&b)
	result := 1
	for i := 1; i < b; i++ {
		result *= a
	}
	fmt.Printf("a power of a number: %d\n", result)
}
