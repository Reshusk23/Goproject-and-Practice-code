package main

import "fmt"

func main() {
	a := 5 //50 replace and try for seeing difrent execution of loop goto statement
LOOP:
	for a < 20 {
		if a == 12 {
			a = a + 1
			goto LOOP
		}
		fmt.Println("value of a", a)
		a++
	}

}
