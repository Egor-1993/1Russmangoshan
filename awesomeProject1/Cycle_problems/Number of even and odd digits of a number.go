package main

import "fmt"

func main() {
	var num int
	fmt.Print("enter number: ")
	fmt.Scan(&num)
	even := 0
	uneven := 0
	for num != 0 {
		digit := num % 10
		if digit%2 == 0 {
			even++
		} else {
			uneven++
		}
		num /= 10
	}
	fmt.Println("sum even:", even)
	fmt.Println("sum uneven: ", uneven)
}
