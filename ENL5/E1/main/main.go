package main

import (
	"fmt"
)

type person struct {
	firstname string
	lastname  string
	ficfs     []string
}

func main() {
	p1 := person{
		firstname: "Some",
		lastname:  "Person",
		ficfs:     []string{"Vanilla", "Strawberry", "Chocolate"},
	}
	p2 := person{
		firstname: "Other",
		lastname:  "Person2",
		ficfs:     []string{"Stracciatella"},
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.firstname, p1.lastname)
	for _, f := range p1.ficfs {
		fmt.Println(f)
	}
	fmt.Println(p2.firstname, p2.lastname)
	for _, f := range p2.ficfs {
		fmt.Println(f)
	}
}
