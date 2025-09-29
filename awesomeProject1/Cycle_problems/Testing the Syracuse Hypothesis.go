package main

import "fmt"

func main() {
	for start := 20; start <= 30; start++ {
		n := start
		fmt.Printf("number %d: ", n)
		for n != 1 {
			if n%2 == 0 {
				n = n / 2
			} else {
				n = (3*n + 1) / 2
			}
			fmt.Printf("%d ", n)

		}
		fmt.Println()
	}
}
