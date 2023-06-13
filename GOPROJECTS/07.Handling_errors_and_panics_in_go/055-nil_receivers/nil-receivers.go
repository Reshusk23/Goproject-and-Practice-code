package main

import "fmt"

type name struct {
	names string
}

func (R *name) str() {
	fmt.Println("hi my name is ", R.names)
}
func main() {
	R := &name{"BOB"}
	R = nil
	R.str()
}
