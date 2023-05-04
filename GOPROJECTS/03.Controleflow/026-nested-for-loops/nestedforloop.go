package main

import "fmt"

func main() {
	/*listnumb := []int{1, 2, 3, 4, 5, 6}
	listalp := []string{"a", "b", "c", "d"}
	for _, i := range listnumb {
		fmt.Println(i)
		for _, l := range listalp {
			fmt.Println(l)
		}
	}*/

	var i, j int
	for i = 2; i < 100; i++ {
		for j = 2; i <= (i / j); j++ {
			if i%j == 0 {
				break
			}
		}
		if j > (i / j) {
			fmt.Printf("%d is prime \n", i)
		}
	}

}
