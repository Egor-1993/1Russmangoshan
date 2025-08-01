package main

import "fmt"

func main() {
	var a int
	fmt.Print("Enter a : ")
	fmt.Scan(&a)
	result := 1
	for i := 2; i <= a; i++ {
		result *= i
	}
	fmt.Printf("a fac: %d\n", result)
}
