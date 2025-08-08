package main

import "fmt"

func main() {
	var nums = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Array in revers order:")
	for i := len(nums) - 1; i >= 0; i-- {
		fmt.Print(nums[i], ",")
	}
}
