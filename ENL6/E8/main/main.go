package main

import (
	"fmt"
)

func main() {
	f := foo()
	fmt.Println(f(1, 2))
}

func foo() func(int, int) int {
	return func(a int, b int) int {
		return a + b
	}
}
