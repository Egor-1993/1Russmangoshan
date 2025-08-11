package main

import "fmt"

func main() {
	nums := []int{51, 21, 22, 38, 65, 45, 87, 59, 56, 722, 444, 484, 26, 654}
	var evens []int
	for _, num := range nums {
		if num%2 == 0 {
			evens = append(evens, num)
		}
	}
	fmt.Println("Even numbers:", evens)
}
