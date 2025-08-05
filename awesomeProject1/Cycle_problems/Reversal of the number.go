package main

import "fmt"

func main() {
	var num int
	fmt.Print("enter number:")
	fmt.Scan(&num)
	reversed := 0
	for num > 0 {
		digit := num % 10
		reversed = reversed*10 + digit
		num /= 10
	}
	fmt.Println("revers number:", reversed)
}
