package main

import "fmt"

func main() {
	var a = 1
	var b = 2
	fmt.Println("a is", a)
	fmt.Println("b is", b)
	fmt.Println("a's address is", &a)
	fmt.Printf("a's address is %d\n", &a)
	fmt.Println("b's address is", &b) // does not work with const
	fmt.Printf("b's address is %d\n", &b)
}
