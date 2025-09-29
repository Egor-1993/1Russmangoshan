package main

import "fmt"

func main() {
	var m, n int
	fmt.Print("enter number m:")
	fmt.Scan(&m)
	fmt.Print("enter number n:")
	fmt.Scan(&n)
	for num := m; num <= n; num++ {
		fmt.Printf("Number divisors-%d: ", num)
		found := false
		for i := 2; i < num; i++ {
			if num%i == 0 {
				fmt.Printf("%d ,", i)
				found = true
			}
		}
		if !found {
			fmt.Print("non. ")
		}
		fmt.Println(" ")
	}
}
