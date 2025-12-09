package main

import "fmt"

func main() {
	var a string
	a = "aceld"

	var allType interface{}
	allType = a

	str, value := allType.(string)
	fmt.Println(str, " _ ", value)
}