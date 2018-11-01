package main

import (
	"fmt"
)

func main() {
	var arr [5]int
	arr[0] = 2
	arr[1] = 4
	arr[2] = 8
	arr[3] = 16
	arr[4] = 32
	for i, v := range arr {
		fmt.Println("Value at index", i, "is", v)
	}
	fmt.Printf("type is %T\n", arr)
}
