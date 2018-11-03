package main

import (
	"fmt"
)

func main() {
	a := store()
	b := store()
	a("x")
	b("y")
	a("z")
}

func store() func(string) {
	var a []string
	return func(s string) {
		a = append(a, s)
		fmt.Println(a)
	}
}
