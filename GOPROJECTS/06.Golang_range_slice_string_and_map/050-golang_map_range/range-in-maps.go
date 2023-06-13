package main

import "fmt"

func main() {
	name:=map[string]int{"bob":50,"david":60,"denny":80,"sammy":20}
	for salary,str:=range name{
		fmt.Println("salary",salary,"====","name",str)
	}
}