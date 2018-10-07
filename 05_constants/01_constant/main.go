package main

import "fmt"

const a string = "Hello"

func main() {
	fmt.Println(a)
	//sa = "Welt"
	const b = a
	fmt.Println(b)
}