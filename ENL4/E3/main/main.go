package main

import (
	"fmt"
)

func main() {
	base := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(base[:5])
	fmt.Println(base[5:])
	fmt.Println(base[2:7])
	fmt.Println(base[1:6])
}
