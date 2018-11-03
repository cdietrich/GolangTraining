package main

import (
	"fmt"
)

func main() {
	f := func(a int, b int) int { return a + b }
	foo(f, 1, 2)
	add := func(a int, b int) int {
		return a + b
	}
	fmt.Println(sum(add, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	minus := func(a int, b int) int {
		return a - b
	}
	fmt.Println(sum(minus, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

func foo(f func(a int, b int) int, a int, b int) {
	fmt.Println(f(a, b))
}

func sum(f func(x int, y int) int, a ...int) int {
	total := 0
	for _, v := range a {
		total = f(total, v)
	}
	return total
}
