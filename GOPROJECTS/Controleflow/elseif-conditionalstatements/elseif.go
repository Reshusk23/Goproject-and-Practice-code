package main

import "fmt"

func main() {
	//if i=2 and i<=2 in this case both are same condition so we use else if statements
	/*x := 50
	if x == 100 {
		fmt.Println("this is 100 value")
	} else if x == 50 {
		fmt.Println("this is 50 value")
	} else {
		fmt.Println("this is last statement")
	}*/

	balance := 50
	if balance < 100 {
		fmt.Println("this is 100 balance")
	} else if balance == 20 {
		fmt.Println("this is 50 balance")
	} else {
		fmt.Println("this is last statement with o")
	}

}
