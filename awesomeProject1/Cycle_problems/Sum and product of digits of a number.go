package main

import "fmt"

func main() {
	var num int
	fmt.Print("enter number: ")
	fmt.Scan(&num)
	sum := 0
	mult := 1
	for num != 0 {
		digit := num % 10
		sum += digit
		mult *= digit
		num /= 10
	}
	fmt.Println("digit sum :", sum)
	fmt.Println("digit mult: ", mult)
}
