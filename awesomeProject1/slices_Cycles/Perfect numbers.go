package main

import (
	"fmt"
)

func main() {
	for num := 1; num <= 10000; num++ {
		sum := 0
		for i := 1; i <= num/2; i++ {
			if num%i == 0 {
				sum += i
			}
		}
		if sum == num {
			fmt.Printf("Perfeckt number:%d\n", num)
		}
	}
}
