package main

import (
	"fmt"
)

func main() {
	p := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	fmt.Println(p)
}
