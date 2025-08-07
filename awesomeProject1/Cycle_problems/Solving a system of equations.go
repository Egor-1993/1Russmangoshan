package main

import "fmt"

func main() {
	for pens := 0; pens <= 30; pens++ {
		for pencils := 0; pencils <= 30; pencils++ {
			erasers := 30 - pens - pencils
			if erasers >= 0 {
				total := pens*10 + pencils*5 + erasers*2
				if total == 100 {
					fmt.Printf("pens: %d, pencils: %d, erasers: %d\n", pens, pencils, erasers)
				}
			}
		}
	}
}
