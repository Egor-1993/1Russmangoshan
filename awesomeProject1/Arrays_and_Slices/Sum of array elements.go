package main

import "fmt"

func main() {
	var nums = [5]int{5, 84, 55, 32, 15}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	fmt.Println("Sum of all array elements: ", sum)
}
