package main

import "fmt"

func main() {
	/*a=[3][4]int{
		{0,1,2,3},
		{4,5,6,7},
		{8,9,10,11}
	}*/

	var a = [6][3]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {9, 10, 11}, {12, 13, 14}, {15, 16, 17}}
	var i, j int
	for i = 0; i < 6; i++ {
		for j = 0; j < 3; j++ {
			fmt.Printf("a[%d][%d]=%d \n", i, j, a[i][j])
		}
	}
}
