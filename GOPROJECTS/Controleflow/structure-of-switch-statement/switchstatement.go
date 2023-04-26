package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()
	switch today.Day() {
	case 1:
		fmt.Println("today is 1st date, this is day of school")
	case 2: //if i put 6 here it will be give output
		fmt.Println("today is 2nd date, this is day of college")
	case 3:
		fmt.Println("today is 3rd date, this is day of university")
	default:
		fmt.Println("No information found")
	}

}
