package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	field := []rune{'@', '@', '@', '@', '@', '@', '@', '@', '@', '@'}
	ship := []int{3, 4, 5, 6}
	hits := make([]bool, len(ship))
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(string(field))
		fmt.Println("Enter a coordinate (0-9):")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err != nil || num < 0 || num > 9 {
			fmt.Println("Invalid coordinate, try again!")
			continue
		}
		hit := false
		for i, pos := range ship {
			if num == pos {
				if hits[i] {
					fmt.Println("You already hit this spot! ")
				} else {
					hits[i] = true
					field[num] = 'X'
					fmt.Println("You hit!")
				}
				hit = true
				break
			}
		}
		if !hit {
			fmt.Println("You missed!")
			field[num] = '.'
		}
		allHit := true
		for _, h := range hits {
			if !h {
				allHit = false
				break
			}
		}
		if allHit {
			fmt.Println("You won!")
			fmt.Println(string(field))
			break
		}
	}
}
