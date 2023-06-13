package main

import "fmt"

func D() {
	fmt.Println("D()")
	defer B()
}

func B() {
	fmt.Println("B()")
	err := recover()
	fmt.Printf("ERROr = %+v", err)
}
func main() {
	D()
	fmt.Println("")
}
