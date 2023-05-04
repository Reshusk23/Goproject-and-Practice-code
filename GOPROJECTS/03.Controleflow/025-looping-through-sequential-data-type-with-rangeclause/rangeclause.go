package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := range numbers {
		fmt.Println("slice item", i, "is", numbers[i])

	}

	fish := []string{"dogfish", "whitefish", "blackfish", "hamerheadfish"}
	for a, fish := range fish {
		fmt.Println(a, fish)
	}
}
