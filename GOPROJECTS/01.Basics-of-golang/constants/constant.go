package main

import (
	"fmt"
	//"math"
)

const j = 3.14

func main() {
	/*fmt.Println("constant")
	var i = math.Sqrt(4)
	fmt.Println("constant", i)*/

	/*const j = math.Sqrt(4)  //floating points never been accepted as constants
	fmt.Println("constant", j)*/

	const i = "hello world" //string constent
	fmt.Println("hello", i)
	fmt.Println("PI", j, "value")

	const t = true
	fmt.Println("bool", t)
}
