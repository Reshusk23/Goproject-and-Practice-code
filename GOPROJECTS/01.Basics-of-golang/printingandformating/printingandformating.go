package main

import "fmt"

var (
	a = 600
	b = false
	c = 3.14
	d = "Go"
)

func main() {
	//printing the variables
	/*fmt.Println("value of a \n", a)
	fmt.Println("value of b \n", b)
	fmt.Println("value of c \n", c)
	fmt.Println("value of d \n", d)*/

	//printing and formating the variables
	fmt.Printf("value of a %d\n", a) // %d is used for formating integers
	fmt.Printf("value of b %t\n", b) // %t is used for formating boolean values
	fmt.Printf("value of c %g\n", c) // %g is used for formating float values
	fmt.Printf("value of d %s\n", d) // %s is used for formating strings
}
