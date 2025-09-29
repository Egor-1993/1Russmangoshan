package main

import "fmt"

func main() {
	nums := []int{65, 15, 464, 78, 52, 32}
	index := 3
	if index < 0 || index >= len(nums) {
		fmt.Println("Index out of range")
		return
	}
	nums = append(nums[:index], nums[index+1:]...)
	fmt.Println("Slice after deletion:", nums)
}
