package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("a")
	fmt.Println("b")
}
