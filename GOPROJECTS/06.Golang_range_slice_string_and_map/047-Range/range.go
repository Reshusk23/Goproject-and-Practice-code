package main

import "fmt"

func main() {
	/*i := []int{1, 2, 3, 4}
	for j := range i {
		i = append(i, j)
		fmt.Println(j)
	}*/

	i := []string{"go", "lang", "python", "java", "c"}
	for j := range i {
		fmt.Println(j)
	}
}
