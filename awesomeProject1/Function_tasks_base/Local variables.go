package main

import "fmt"

func scopedVar() int {
	x := 5
	return x // ← вот теперь функция "отдаёт" x наружу
}

func main() {
	result := scopedVar()
	fmt.Println("x:", result)
}
