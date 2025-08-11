package main

import "fmt"

func main() {
	nums := []int{15, 25, 35, 45, 55}
	first := nums[0]
	for i := 0; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
	}
	nums[len(nums)-1] = first
	fmt.Println("After shift left:", nums)
}
