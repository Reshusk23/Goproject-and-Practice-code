package main

import "fmt"

func main() {
	/*if stst{
	fmt.Println("true")
	if nestedstst{
		fmt.Println("yes")
	}else{
		fmt.Println("no")
	}else{
		fmt.Println("false")
	}*/

	x := 50
	y := 100
	if x == 50 { //if i set here 100 then no out put gets
		if y == 50 { //put 100 insted of 50 here for your understanding
			fmt.Printf("value of x is 50 and y is 100\n")
		} else {
			fmt.Println("this is inner else statement")
		}
	} else {
		fmt.Println("value is false")
	}
}
