package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {

			fmt.Println("breaking the loop")
			break
		}
		fmt.Println("the value of i is", i)
	}

	fmt.Println("exiting program")
}
