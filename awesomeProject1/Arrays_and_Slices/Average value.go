package main

import "fmt"

func main() {
	var nums = [10]float64{10, 20, 30, 50, 40, 90, 80, 70, 100, 60}
	sum := 0.0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	average := sum / float64(len(nums))
	fmt.Println("Average value of array:", average)
}
