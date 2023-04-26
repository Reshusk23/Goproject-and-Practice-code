package main

import "fmt"

func main() {
	p := 20
	d := 15
	fmt.Println(1 + 6)
	fmt.Println(1 - 6)
	fmt.Println(1 * 6)
	fmt.Println(1 / 6)
	fmt.Println(1 % 6)
	c := p + d
	fmt.Printf("Result of addition is= %d \n", c)
	c1 := p - d
	fmt.Printf("Result of subtraction is= %d \n", c1)
	c2 := p * d
	fmt.Printf("Result of multiplication is= %d \n", c2)
	c3 := p / d
	fmt.Printf("Result of divistion is= %d \n", c3)
	c4 := p % d
	fmt.Printf("Result of modulardivition is= %d \n", c4)
}
