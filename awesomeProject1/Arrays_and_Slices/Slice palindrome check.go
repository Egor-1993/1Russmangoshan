package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter number separated by space:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	parts := strings.Split(input, " ")
	nums := make([]int, 0, len(parts))
	for _, p := range parts {
		num, err := strconv.Atoi(p)
		if err != nil {
			fmt.Println("Invalid input :", p)
			return
		}
		nums = append(nums, num)
	}
	isPalindrome := true
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		if nums[i] != nums[j] {
			isPalindrome = false
			break
		}

	}
	fmt.Printf("Input slice:%v\n", nums)
	fmt.Printf("Is palindrome:%v\n", isPalindrome)
}
