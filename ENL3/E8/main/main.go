package main

import (
	"fmt"
)

func main() {
	a := 0
	switch {
	case a == 1:
		fmt.Println("one")
	case a == 0:
		fmt.Println("zero")
	default:
		fmt.Println("default")
	}
}
