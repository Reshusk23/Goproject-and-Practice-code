package main

import "fmt"

func main() {
	a := 4
	b := 5
	c1 := a == b
	fmt.Println(c1)
	c2 := a != b
	fmt.Println(c2)
	c3 := a < b
	fmt.Println(c3)
	c4 := a <= b
	fmt.Println(c4)
	c5 := a > b
	fmt.Println(c5)
	c6 := a >= b
	fmt.Println(c6)
}
