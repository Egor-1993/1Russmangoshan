package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Enter a: ")
	fmt.Scan(&a)
	fmt.Print("Enter b: ")
	fmt.Scan(&b)
	for i := a; i <= b; i++ {
		fmt.Printf("%d ", i*i*i)

	}
}
