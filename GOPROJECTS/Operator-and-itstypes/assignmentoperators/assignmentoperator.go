package main

import "fmt"

func main() {
	a := 20
	a += 10
	fmt.Println(a)
	var b int = 50
	var c int = 20
	b = c
	fmt.Println(b)
	b += c
	fmt.Println(b)
	b -= c
	fmt.Println(b)
	b /= c
	fmt.Println(b)
	b %= c
	fmt.Println(b)
	b *= c
	fmt.Println(b)
}
