package main

import "fmt"

func main() {
	add(46, 5)
	add(255, 6)
}

func add(x int, y int) {
	var z = x + y
	fmt.Println("x + y = ", z)
}
