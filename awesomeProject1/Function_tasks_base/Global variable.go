package main

import "fmt"

var counter int = 0

func increment() {
	counter = counter + 1
}
func main() {
	increment()
	increment()
	increment()
	increment()
	increment()
	fmt.Println("Counter:", counter)
}
