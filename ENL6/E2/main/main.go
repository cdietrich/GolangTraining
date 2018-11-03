package main

import (
	"fmt"
)

func main() {
	fmt.Println(foo([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}...))
	fmt.Println(bar([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}

func foo(a ...int) int {
	result := 0
	for i := 0; i < len(a); i++ {
		result += a[i]
	}
	return result
}

func bar(a []int) int {
	return foo(a...)
}
