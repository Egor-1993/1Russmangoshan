package main

import "fmt"

func countdown(n int) {
	if n == 0 {
		return
	}
	fmt.Print(n, " ")
	countdown(n - 1)
}
func main() {
	var number int
	fmt.Print("Enter your number: ")
	fmt.Scan(&number)
	countdown(number)
}
