package main

import (
	"fmt"
)

func main() {
	x := struct {
		a string
		b string
	}{a: "A", b: "B"}
	fmt.Println(x)
	y := struct {
		a string
		b string
	}{"a", "b"} // discouraged
	fmt.Println(y)
}
