package main

import (
	"fmt"
)

func main() {
	var side int
	fmt.Print("Введите сторону квадрата: ")
	fmt.Scan(&side)

	area := side * side
	fmt.Println("Площадь квадрата:", area)
}
