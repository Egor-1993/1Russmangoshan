package main

import "fmt"

func main() {
	var num int
	fmt.Print("enter number: ")
	fmt.Scan(&num)
	if num < 0 {
		num = -num
	}
	sum := 0
	for num > 0 {
		digit := num % 10
		if digit%2 == 0 {
			sum += digit
		}
		num /= 10
	}
	fmt.Println("Sum of even digits: ", sum)
}
