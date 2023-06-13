package main

import "fmt"

func D(n1, n2 int) int {
	defer func() {
		fmt.Println(recover())
	}()
	q := n1 / n2
	return q
}
func main() {
	fmt.Println(D(50, 60))
	fmt.Println(D(50, 0))
}
