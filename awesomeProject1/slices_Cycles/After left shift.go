package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("original slices:", nums)
	first := nums[0]
	nums = nums[1:]
	nums = append(nums, first)
	fmt.Println("after left shift:", nums)
}
