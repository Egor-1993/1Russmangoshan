package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a positive integral: ")
	fmt.Scan(&num)
	lastDigit := num % 10
	firstDigit := num
	for firstDigit >= 10 {
		firstDigit /= 10
	}
	sum := firstDigit + lastDigit
	fmt.Printf("Sum of first and last digits:%d\n", sum)
}
