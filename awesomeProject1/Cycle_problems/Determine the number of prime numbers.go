package main

import "fmt"

func main() {
	var count, input int
	fmt.Println("Enter integers (completion - any number < 2):")
	for {
		fmt.Print("enter number:")
		fmt.Scan(&input)
		if input < 2 {
			break
		}
		isPrime := true
		for i := 2; i*i <= input; i++ {
			if input%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			count++
		}
	}
	fmt.Printf("Number of prime numbers: %d\n", count)
}
