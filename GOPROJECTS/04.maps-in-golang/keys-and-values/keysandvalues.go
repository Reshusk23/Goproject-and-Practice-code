package main

import "fmt"

func main() {
	/*var match = make(map[string]string)

	match["Bob"] = "Sana"
	match["david"] = "Sophia"
	fmt.Println(match)
	match["bob"] = "Emma"
	fmt.Println(match)*/

	key := []string{}
	for keys := range bob {
		key = append(key, keys)
	}
	fmt.Printf("%q", key)
}
