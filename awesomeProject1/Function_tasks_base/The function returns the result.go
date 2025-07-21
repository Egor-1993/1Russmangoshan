package main

import "fmt"

func square(n int) int {
	return n * n

}
func main() {
	var x, y int
	fmt.Print("Enter x: ")
	fmt.Scan(&x)
	fmt.Print("Enter y: ")
	fmt.Scan(&y)
	result := square(x) + square(y)
	fmt.Println("Square:", result)
}
