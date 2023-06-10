package main

import "fmt"

func main() {
	color := []string{"Red", "green", "black", "pink", "white"}
	color = append(color, "black-white")
	color = append(color, "black-red", "red-blue")
	fmt.Printf("%q\n", color)
	color = append(color[:2], color[6:]...)
	fmt.Printf("%q \n", color)
}
