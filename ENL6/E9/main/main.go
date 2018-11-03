package main

import (
	"fmt"
)

func main() {
	doit(func() int {
		return 42
	})
}

func doit(a func() int) {
	fmt.Println(a())
}
