package main

import "fmt"

func main() {
	/*var a map[string]int //any map key and values are nill then it shows error in the runtime
	fmt.Println(a)
	if a == nil {
		fmt.Println("your A value is nill")
	}*/

	Animal := map[string]string{"name": "cow", "color": "black", "bread": "astralian"}
	fmt.Println(Animal)
	fmt.Println(Animal["name"])
	fmt.Println(Animal["color"])
}
