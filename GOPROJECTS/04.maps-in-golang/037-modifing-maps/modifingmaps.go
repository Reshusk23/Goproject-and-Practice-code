package main

import "fmt"

func main() {

	/*countrymaps := map[string]string{"USA": "W-DC", "UK": "London", "Germany": "Berlin"}
	countrymaps["USA"] = "New york"
	fmt.Println("this is updated map", countrymaps)*/

	countrymaps := map[string]int{"USA": 500, "UK": 400, "Germany": 600}
	countrymaps["USA"] = 900
	fmt.Println(countrymaps)
}
