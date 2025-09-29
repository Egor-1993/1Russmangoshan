package main

import "fmt"

func main() {
	var nums = [8]int{21, 51, 54, 55, 68, 998, 12, 34}
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			count++
		}
	}
	fmt.Println("Number of even elements in array:", count)
}
