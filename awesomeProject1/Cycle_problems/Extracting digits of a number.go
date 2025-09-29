package main

import "fmt"

func main() {
	var num int
	fmt.Print("enter number: ")
	fmt.Scan(&num)
	if num < 0 {
		num = -num

	}
	fmt.Println("digit: ")
	for num > 0 {
		digit := num % 10
		fmt.Println(digit)
		num = num / 10
	}
}
