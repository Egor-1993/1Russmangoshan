package main

import "fmt"

func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}

}
func main() {
	double := makeMultiplier(2)
	var num int
	fmt.Print("Enter number: ")
	fmt.Scan(&num)
	result := double(num)
	fmt.Println("double number:", result)
}
