package main

import "fmt"

func main() {
	for i := 32; i <= 126; i++ {
		fmt.Printf("%3d:%c | ", i, i)
		if (i-31)%5 == 0 {
			fmt.Println()

		}
	}
}
