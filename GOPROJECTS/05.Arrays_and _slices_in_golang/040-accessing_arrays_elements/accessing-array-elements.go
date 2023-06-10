package main

import "fmt"

func main() {
	var n [10]int
	var a, b int

	for a = 0; a < 10; a++ {
		n[a] = a + 100
	}
	for b = 0; b < 10; b++ {
		fmt.Printf("Element[%d]=%d\n", b, n[b])
	}
}
