package main

import "fmt"

func main() {
	nums := []int{15, 11, 2121, 32, 54}
	sub := nums[1:4]
	fmt.Println("Original slice:", nums)
	fmt.Println("Sub slice:", sub)
	sub[1] = 500
	fmt.Println("\nAfter modification:")
	fmt.Println("Original slice:", nums)
	fmt.Println("Sub slice:", sub)
}
