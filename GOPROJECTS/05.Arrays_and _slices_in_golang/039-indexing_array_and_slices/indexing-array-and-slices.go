package main

import "fmt"

func main() {
	numbers := [3]string{"One", "Two", "Three"}

	fmt.Printf("%q\n", numbers)
	fmt.Println(numbers[1])
	fmt.Println(numbers[2])
	fmt.Println(numbers[0])
}
