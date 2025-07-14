package main

import (
	"fmt"
)

func main() {
	var celsius float64
	fmt.Print("Enter temperature in Celsius: ")
	fmt.Scan(&celsius)
	Fahrenheit := celsius*1.8 + 32
	fmt.Print(Fahrenheit)
}
