package main

import (
	"fmt"
)

func main() {
	var a [5]int
	fmt.Println(a)
	a[3] = 1
	fmt.Println(a)
	fmt.Println(len(a))
	for i, ai := range a {
		fmt.Println(ai, "at index", i)
	}
}
