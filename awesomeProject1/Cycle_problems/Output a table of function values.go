package main

import (
	"fmt"
)

func main() {
	for x := -5.0; x <= 5.0; x += 0.5 {
		y := 5 - (x * x / 2)
		fmt.Printf("x=%.1f | y=%.2f\n", x, y)

	}
}
