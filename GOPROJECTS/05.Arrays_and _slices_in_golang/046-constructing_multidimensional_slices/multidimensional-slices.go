package main

import "fmt"

func main() {
	s1 := [][]int{{10, 15}, {20, 30}, {50}, {60, 70}}
	for _, x := range s1 {
		for _, y := range x {
			fmt.Println(y)
		}
	}
	fmt.Println(s1)
	fmt.Println(s1[0])
	fmt.Println(s1[0][1])
	fmt.Println(s1[0][0])
}
