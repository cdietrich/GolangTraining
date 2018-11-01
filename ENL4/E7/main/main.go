package main

import (
	"fmt"
)

func main() {
	data := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}
	fmt.Println(data)
	for i, s := range data {
		fmt.Println("Entry", i)
		for _, v := range s {
			fmt.Println(v)
		}
	}
}
