package main

import "fmt"

func main() {
	/*x := 50
	if x == 50{		//50
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}*/

	grade := 80
	if grade <= 60 { //changed the > symbol in this
		fmt.Println("pass the exam")
	} else {
		fmt.Println("fail the exam")
	}
}
