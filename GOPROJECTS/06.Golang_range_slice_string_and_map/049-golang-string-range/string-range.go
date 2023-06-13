package main

import "fmt"

func main() {
	s1 := []string{"go", "lang", "python", "java"}
	s2 := "programing"
	for i := range s1 {
		fmt.Println(i)
	}
	for i, j := range s2 {
		fmt.Println(i, string(j))
	}
}
