package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

func changeMe(p *person, first string, last string) {
	p.first = first
	(*p).last = last
}

func dontChangeMe(p person, first string, last string) {
	p.first = first
	p.last = last
}

func main() {
	p := person{
		first: "Miss",
		last:  "Moneypenny",
	}
	fmt.Println(p)
	changeMe(&p, "Mrs", "Bond")
	fmt.Println(p)
	dontChangeMe(p, "Dr", "No")
	fmt.Println(p)
}
