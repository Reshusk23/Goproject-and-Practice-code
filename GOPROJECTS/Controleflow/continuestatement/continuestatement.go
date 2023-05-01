package main

import "fmt"

func main() {
	/*a := 10
	for a < 20 {
		if a == 15 {
			a = a + 1
			continue
		}
		fmt.Println("value of a", a)
		a++
	}*/

	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("continuing loop")
			continue
		}
		fmt.Println("the value of i", i)
	}
	fmt.Println("exit program")
}
