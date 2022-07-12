package main

import "fmt"

func getType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int: ", v)
	case string:
		fmt.Println("string: ", v)
	case bool:
		fmt.Println("bool: ", v)
	case chan interface{}:
		fmt.Println("chan interface{}: ", v)
	default:
		fmt.Println("I don't know this type: ", v)
	}
}
func main() {
	channel := make(chan interface{})

	getType(10)
	getType("hello world")
	getType(true)
	getType(channel)
}
