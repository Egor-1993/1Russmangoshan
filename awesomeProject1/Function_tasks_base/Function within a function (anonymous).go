package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Enter a: ")
	fmt.Scan(&a)
	fmt.Print("Enter b: ")
	fmt.Scan(&b)
	add := func(a, b int) int {
		return a + b
	}
	result := add(a, b)
	fmt.Println("sum a,b:", result)
}
