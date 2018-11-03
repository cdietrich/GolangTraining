package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	p := person{first: "Miss", last: "Moneypenny"}
	fmt.Println(p.full())
	s := secretAgent{person: person{first: "James", last: "Bond"}, ltk: true}
	fmt.Println(s.full())
}

func (p person) full() string {
	return fmt.Sprint(p.first, " ", p.last)
}
