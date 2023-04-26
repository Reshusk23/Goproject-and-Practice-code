package main

import "fmt"

func main() {
	var i int8 = 10
	var j int32
	j = int32(i)
	fmt.Println(j)
	var f float64 = 200.56
	var d int = int(f)
	fmt.Printf("f=%.2f\n", f)
	fmt.Printf("d=%d\n", d)
}
