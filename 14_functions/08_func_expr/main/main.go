package main

import (
	"fmt"
)

func main() {
	add := func(a int, b int) int {
		return a + b
	}
	fmt.Println(add(1, 2))
	fmt.Println(add(3, 4))
	fmt.Println(add(5, 6))
}
