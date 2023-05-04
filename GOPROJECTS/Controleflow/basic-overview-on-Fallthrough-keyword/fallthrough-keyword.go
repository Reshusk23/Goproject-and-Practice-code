package main

import "fmt"

func main() {
	burger := []string{"beef burger", "turkey burger", "bean burger"}
	for _, favorite := range burger {
		switch favorite {
		case "beef burger":
			fmt.Println(favorite, "is my favorite")
			fallthrough

		case "turkey burger", "bean burger":
			fmt.Println(favorite, "these are greate of taste")

		default:
			fmt.Println("i had never taste this", favorite, "before")
		}
	}
}
