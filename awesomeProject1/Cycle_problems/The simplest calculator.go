package main

import "fmt"

func main() {
	var x, y float64
	var op string
	var result float64
	for {
		fmt.Print("Enter the first number:")
		fmt.Scan(&x)
		fmt.Print("Enter an operator(+,-,*,/) or 0 to exit:")
		fmt.Scan(&op)
		if op == "0" {
			fmt.Printf("Exiting the calculator. Bye! ")
			break
		}
		fmt.Print("Enter the second number: ")
		fmt.Scan(&y)
		switch op {
		case "+":

			result = x + y
		case "-":
			result = x - y
		case "*":
			result = x * y
		case "/":
			if y == 0 {
				fmt.Println("Error: division by zero is not allowed.")
				continue
			}
			result = x / y
		default:
			fmt.Println("Invalid operator. Please use +, -, *, / or 0 to exit.")
			continue
		}
		fmt.Printf("Result: %.2f\n", result)
	}
}
