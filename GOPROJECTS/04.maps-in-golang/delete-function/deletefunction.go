package main

import "fmt"

func main() {

	countrymaps := map[string]string{"USA": "W-DC", "UK": "London", "Germany": "Berlin"}
	fmt.Println("this is my original map")
	for country := range countrymaps {
		fmt.Println("capital maps", country, "is", countrymaps[country])
	}
	delete(countrymaps, "USA")
	fmt.Println("Entry for USA is deleted")
	fmt.Println("This map updated")
	for country := range countrymaps {
		fmt.Println("capital maps of", country, "are", countrymaps[country])
	}

}
