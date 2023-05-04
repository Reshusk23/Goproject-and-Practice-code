package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "this is a string package"
	fmt.Println(strings.ToUpper(a))
	b := "THIS IS A STRING PACKAGE UPPER CASE"
	fmt.Println(strings.ToLower(b))
}
