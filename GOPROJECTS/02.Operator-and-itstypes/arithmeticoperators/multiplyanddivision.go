package main

import "fmt"

func main() {
	a := 50.0 //only 2 integers and 2 floating points do multiplication if one variable is int and another one is float then that will show ous error
	b := 20.0
	c := float64(a) / float64(b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(c)

}
