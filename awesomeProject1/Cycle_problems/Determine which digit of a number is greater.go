package main

import "fmt"

func main() {
	var num int
	fmt.Print("enter number")
	fmt.Scan(&num)
	maxDigit := 0
	for num != 0 {
		digit := num % 10
		if digit > maxDigit {
			maxDigit = digit
		}
		num /= 10
	}
	fmt.Println("Max digit:", maxDigit)
}
