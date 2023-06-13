package main

import (
	"errors"
	"fmt"
)

func main() {
	error1 := errors.New("ERROR USER GENERATED")
	fmt.Println(error1)
}
