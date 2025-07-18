package main

import "fmt"

func double(n int) int {
	var result = n * 2
	return result
}
func main() {
	fmt.Println(double(5))
	fmt.Println(double(-651))
	fmt.Println(double(7))
	fmt.Println(double(889))
	fmt.Println(double(965))
	fmt.Println(double(-5))
}
