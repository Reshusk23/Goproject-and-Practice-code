package main

import "fmt"

func main() {
	//for[int stst];[condition];[post stst]
	//block of code

	for i := 0; i < 5; i++ {
		if i != 5 {
			fmt.Printf("welcome %d to loop \n", i)
		}
		fmt.Println("value of i is", i)
	}
}
