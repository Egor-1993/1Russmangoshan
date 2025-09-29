package main

import "fmt"

func main() {
	var nums = [10]int{51, 65, 7987, 51, 32, 45, 987, 154, 522, 979}
	target := 987
	found := false
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			fmt.Printf("Element found at index: %d (position: %d)\n", i, i+1)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Element not found in array:")
	}
}
