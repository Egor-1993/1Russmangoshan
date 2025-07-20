package main

import "fmt"

func length(s string) int {
	return len(s)
}
func main() {
	var input string
	fmt.Print("Enter text: ")
	fmt.Scan(&input)
	result := length(input)
	fmt.Println(result)
}
