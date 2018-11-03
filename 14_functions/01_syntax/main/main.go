package main

import (
	"fmt"
)

func main() {
	foo()
	bar("World")
	x := plus(1, 2)
	fmt.Println(x)
	d, m := divandmod(5, 2)
	fmt.Println(d, m)
	n := 1
	testPassByValue(n)
	// no change
	fmt.Println(n)
}

// func (r receiver) name(params) (returs) {...}

func foo() {
	fmt.Println("foo")
}

func bar(s string) {
	fmt.Println("Hello", s)
}

func plus(a int, b int) int {
	return a + b
}

func divandmod(a int, b int) (int, int) {
	x := a / b
	y := a % b
	return x, y
}

func testPassByValue(n int) {
	n = n + 1
}
