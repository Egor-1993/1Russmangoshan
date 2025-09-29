package main

import "fmt"

func main() {
	var num int
	fmt.Print("enter number:")
	fmt.Scan(&num)
	lastDigit := num % 10
	firstDigit := num
	for firstDigit >= 10 {
		firstDigit /= 10
	}
	sum := lastDigit + firstDigit
	fmt.Println("sum last and first digit:", sum)
}
