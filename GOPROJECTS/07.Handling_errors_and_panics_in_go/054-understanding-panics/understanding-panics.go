package main

import (
	"os"
)

func main() {
	panic("this is panic")
	_, er := os.Open("/tmp/file")
	if er != nil {
		panic(er)
	}
}
