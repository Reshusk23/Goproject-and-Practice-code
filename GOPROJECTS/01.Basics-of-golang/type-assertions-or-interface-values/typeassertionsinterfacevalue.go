package main

import "fmt"

func main() {
	var i interface{} = "3.14" //"Go" //firstly that will get the primary assertion ie., here i
	s := i.(string)            //primary exprestions.signe and a (type)-"primary exprestion.(type)"
	fmt.Println(s)
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)
	f = i.(float64) //panic error y because it is not a string its a loat and i dont hold or write any interface values
	fmt.Println(f)
}
