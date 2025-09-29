package main

import "fmt"

func isEven(n int) bool {
	return n%2 == 0
}
func main() {
	var number int
	fmt.Println("Enter a number: ")
	fmt.Scan(&number)
	result := isEven(number)
	fmt.Println(result)
}
