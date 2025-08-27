package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter a natural number:")
	fmt.Scan(&n)
	sum := 0
	for n > 0 {
		digit := n % 10
		if digit%2 == 0 {
			sum += digit
		}
		n /= 10
	}
	fmt.Printf("Sum of even digits:%d\n", sum)
}
