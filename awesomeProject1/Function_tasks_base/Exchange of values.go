package main

import "fmt"

func swap(x, y int) (int, int) {
	return y, x
}
func main() {
	var a, b int
	fmt.Println("enter a: ")
	fmt.Scan(&a)
	fmt.Println("enter b: ")
	fmt.Scan(&b)
	fmt.Println("before a=", a, "b=", b)
	a, b = swap(a, b)
	fmt.Println("after a=", a, "b=", b)
}
