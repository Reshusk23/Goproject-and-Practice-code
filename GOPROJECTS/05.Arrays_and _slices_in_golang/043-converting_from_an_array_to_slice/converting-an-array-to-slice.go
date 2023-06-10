package main

import "fmt"

func main() {
	color := []string{"Red", "green", "black", "pink", "white"}
	color = append(color, "black-white")
	color = append(color, "black-red", "red-blue")
	fmt.Printf("%q \n", color)
}
