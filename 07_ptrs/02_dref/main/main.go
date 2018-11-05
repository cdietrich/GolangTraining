package main

import "fmt"

func main() {
	a := 42
	b := &a
	fmt.Println(a, "at", &a, "==", *b, "at", b)

	c := 42
	fmt.Println(&c, *&c)
}
