package main

import "fmt"

func main() {
	s1 := []string{"go", "java", "python", "c++"}
	for i, j := range s1 {
		fmt.Printf("i=%d and element =%s \n", i+1, j)
	}
}
