package main

import (
	"fmt"
)

func main() {
	var money float64
	fmt.Print("Enter Money: ")
	fmt.Scan(&money)
	netmoney := money * 0.87
	{
		fmt.Printf("%.2f\n", netmoney)
	}
}
