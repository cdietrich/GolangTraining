package main

import "fmt"

func main() {
	doit(1)
	doit("Hello")
	doit(false)
}

func doit(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("it's an int")
	case string:
		fmt.Println("it's a string")
	default:
		fmt.Println("don't know")
	}
}
