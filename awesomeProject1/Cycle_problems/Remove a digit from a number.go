package main

import "fmt"

func main() {
	var num, digitToDelete, digit int
	fmt.Print("enter number:")
	fmt.Scan(&num)
	fmt.Print("enter digit to delete:")
	fmt.Scan(&digitToDelete)
	newNum := 0
	multiplier := 1
	for num > 0 {
		digit = num % 10
		if digit != digitToDelete {
			newNum = digit*multiplier + newNum
			multiplier *= 10
		}
		num /= 10
	}
	fmt.Printf("new number:%d\n", newNum)
}
