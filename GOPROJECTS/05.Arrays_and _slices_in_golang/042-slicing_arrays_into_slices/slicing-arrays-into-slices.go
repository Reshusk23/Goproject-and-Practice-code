package main

import "fmt"

func main() {
	color := [5]string{"Red", "green", "black", "pink", "white"}
	fmt.Println(color[1:3])
	fmt.Println(color[1:(1 + 2)])
}
