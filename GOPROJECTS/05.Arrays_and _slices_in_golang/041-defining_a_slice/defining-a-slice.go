package main

import "fmt"

func main() {
	/*animal := []string{"Cow", "Fish", "Bull"}
	fmt.Printf("%q \n", animal)
	anim := make([]string, 2)
	fmt.Println(anim)*/

	var animal []int
	animal = []int{5, 6, 5, 8, 9, 2, 3} //put ==
	animal = make([]int, 5, 6)
	fmt.Println(animal)
}
