package main

import "fmt"

func main() {
	var i interface{} = "Howdy fellow Gophers"

	fmt.Println("Type asserting test string:")
	s, ok := i.(string)
	fmt.Println(s, ok)
	fmt.Println()

	fmt.Println("Type asserting test float64:")
	f, ok := i.(float64)
	fmt.Println(f, ok)
	fmt.Println()

	fmt.Println("Type asserting byte slice:")
	b, ok := i.([]byte)
	fmt.Println(b, ok)
}
