package main

import "fmt"

func main() {
	if a := 1; true {
		fmt.Println(a)
	}
	// this is not valid cause a is no longer in scope
	// fmt.Println(a)
}
