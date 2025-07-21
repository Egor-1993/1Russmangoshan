package main

import "fmt"

var name string = "Egor"

func changeName() {
	name = "Kirill"
}
func main() {
	fmt.Println("before", name)
	changeName()
	fmt.Println("after", name)

}
