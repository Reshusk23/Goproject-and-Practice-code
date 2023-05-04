package main

import "fmt"

func main() {
	var emp = make(map[string]int)
	emp["Bob"] = 20
	emp["devid"] = 30
	fmt.Println(emp)

	emplist := make(map[string]int)
	emplist["Bob"] = 20
	emplist["devid"] = 30
	fmt.Println(emplist)
}
