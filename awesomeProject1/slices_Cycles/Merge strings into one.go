package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"Merge", "strings", "into", "one"}
	result := strings.Join(words, " ")
	fmt.Println("Slices:", words)
	fmt.Println("Joined string:", result)
}
