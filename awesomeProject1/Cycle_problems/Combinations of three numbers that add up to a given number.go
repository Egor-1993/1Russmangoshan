package main

import "fmt"

func main() {
	var limit, target int
	fmt.Print("Enter value limit:")
	fmt.Scan(&limit)
	fmt.Print("Enter amount:")
	fmt.Scan(&target)
	for i := 0; i <= limit; i++ {
		for j := 0; j <= limit; j++ {
			for k := 0; k <= limit; k++ {
				if i+j+k == target {
					fmt.Printf("Combinations of three numbers: %d + %d + %d = %d\n ", i, j, k, target)
				}

			}
		}
	}
}
