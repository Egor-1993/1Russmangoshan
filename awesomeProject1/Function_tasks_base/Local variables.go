package main

import "fmt"

func scopedVar() {
	x := 5
	fmt.Println("x:", x)
}
func main() {
	scopedVar()
	fmt.Print(x)
}
