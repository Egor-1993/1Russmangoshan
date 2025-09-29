package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"Go", "is", "GOYDA"}
	result := strings.Join(words, " ")
	fmt.Println(result)
}
