package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sumit(s...))
}

func sumit(a ...int) int {
	result := 0
	for _, v := range a {
		result += v
	}
	return result
}
