package main

import "fmt"

func main() {
	var nums = [10]int{21, 651, 454, 87, 231, 465, 879, 5464, 321, 155}
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	fmt.Println("Maximum element in array:", max)
}
