package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Enter the number of elements in the series:")
	fmt.Scan(&n)
	sum := 0.0
	correntTerm := 1.0
	for i := 1; i <= n; i++ {
		sum += correntTerm
		correntTerm *= -0.5
	}
	fmt.Printf("Sum of first %d elements of series: %f\n", n, sum)
}
