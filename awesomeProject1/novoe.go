package main

import "fmt"

func main() {
	var num int
	fmt.Print("введите число")
	fmt.Scan(&num)
	if num == 0 {
		fmt.Print("число равно нулю")
	} else {
		fmt.Print("число не равно нулю")
	}
}
