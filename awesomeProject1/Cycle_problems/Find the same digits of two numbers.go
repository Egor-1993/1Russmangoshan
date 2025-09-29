package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Print("enter first number:")
	fmt.Scan(&num1)
	fmt.Print("enter the second number:")
	fmt.Scan(&num2)
	var digits1 [10]bool
	var digits2 [10]bool
	for num1 > 0 {
		d := num1 % 10
		digits1[d] = true
		num1 /= 10
	}
	for num2 > 0 {
		d := num2 % 10
		digits2[d] = true
		num2 /= 10
	}
	fmt.Print("same numbers:")
	found := false
	for i := 0; i <= 9; i++ {
		if digits1[i] && digits2[i] {
			fmt.Print(i, ",")
			found = true
		}
	}
	if !found {
		fmt.Print("no identical numbers.")
	}
}
