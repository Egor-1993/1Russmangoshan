package main

import (
	"fmt"
	"os"
)

func main() {
	var num int
	fmt.Fprint(os.Stdout, "Enter number: ")
	fmt.Scan(&num)
	lasTwo := num % 100
	fmt.Print("Last two digits", lasTwo)
}
