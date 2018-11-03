package main

import (
	"fmt"
)

func main() {
	fmt.Println(sumit(2, 3))
	fmt.Println(sumit())
}

func sumit(a ...int) int {
	fmt.Println(a == nil)
	fmt.Printf("%v,\n%T\n", a, a)
	result := 0
	for _, v := range a {
		result += v
	}
	return result
}
