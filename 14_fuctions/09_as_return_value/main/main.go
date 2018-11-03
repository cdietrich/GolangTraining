package main

import (
	"fmt"
)

func main() {
	f := foo()
	fmt.Printf("%T\n", f)
	fmt.Println(f(1, 2))
}

func foo() func(a int, b int) int {
	return bar
}

func bar(a int, b int) int {
	return a + b
}
