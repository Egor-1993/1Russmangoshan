package main

import "fmt"

func main() {
	var nums = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	reversed := make([]int, 0, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		reversed = append(reversed, nums[i])
	}
	fmt.Println("Reversed slice:", reversed)
}
