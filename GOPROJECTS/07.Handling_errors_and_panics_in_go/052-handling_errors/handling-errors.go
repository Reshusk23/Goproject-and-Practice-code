// custom error generating
package main

import "fmt"

type MyError struct{}

func (a *MyError) Error() string {
	return "errorssssssssssss"
}
func error1() (string, error) {
	return "", &MyError{}
}
func main() {
	e1, er := error1()
	if er != nil {
		fmt.Println("error generated", er)
	}
	fmt.Println("string error", e1)
}
