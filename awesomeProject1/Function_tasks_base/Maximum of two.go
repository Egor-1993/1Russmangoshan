package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}

}
func main() {
	var a, b int
	fmt.Print("Enter a: ")
	fmt.Scan(&a)
	fmt.Print("Enter b: ")
	fmt.Scan(&b)
	result := max(a, b)
	fmt.Println("Maximum of two:", result)
}
