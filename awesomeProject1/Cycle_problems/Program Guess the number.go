package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	var guess int
	fmt.Println("Guess the number(from 1 to 100)!")
	for {
		fmt.Print("Enter yor guess: ")
		fmt.Scan(&guess)
		if guess > target {
			fmt.Println("Too high!")
		} else if guess < target {
			fmt.Println("Too low!")
		} else {
			fmt.Println("Correct! You guessed the number!")
			break
		}
	}
}
