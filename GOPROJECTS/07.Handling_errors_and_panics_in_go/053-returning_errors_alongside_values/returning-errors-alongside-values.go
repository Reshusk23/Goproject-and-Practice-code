package main

import (
	"errors"
	"fmt"
	"strings"
)

func hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("error of new name")
	}
	return strings.ToTitle(name), nil
}
func main() {
	str, er := hello("BOB")
	if er != nil {
		fmt.Println("cloud not say hello", er)
		return
	}
	fmt.Println("Name", str)
}
