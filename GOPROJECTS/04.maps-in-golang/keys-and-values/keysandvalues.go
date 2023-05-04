package main

import "fmt"

func main() {
	/*var match = make(map[string]string)

	match["Bob"] = "Sana"
	match["david"] = "Sophia"
	fmt.Println(match)
	match["bob"] = "Emma"
	fmt.Println(match)*/

	/*bob := map[string]string{
		"Bob":   "Sana",
		"david": "Sophia",
		"bob":   "Emma",
	}

	key := make([]string, 0)
	for keys := range bob {
		key = append(key, keys)
	}
	fmt.Printf("%q", key)*/

	bob := map[string]int{
		"foo": 1,
		"bar": 2,
		"baz": 3,
	}

	key := []string{}
	for keys := range bob {
		key = append(key, keys)
	}
	fmt.Printf("%q", key)
}
