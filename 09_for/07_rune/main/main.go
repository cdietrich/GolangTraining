package main

import "fmt"

func main() {
	var a rune = 2
	var b rune = 'ä'
	fmt.Println(a)
	fmt.Println(b)

	var myString = "Hello, 世界"
	fmt.Println(myString)
	var slice = []byte(myString)
	for c := range slice {
		fmt.Println(c)
	}

	for i := 50; i <= 140; i++ {
		fmt.Println(i, string(i), []byte(string(i)))
	}
	for i := 50; i <= 140; i++ {
		fmt.Printf("%v %v %v\n", i, string(i), []byte(string(i)))
	}
}
