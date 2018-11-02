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
	m := map[string]person{
		p1.lastname: p1,
		p2.lastname: p2,
	}
	for _, v := range m {
		fmt.Println(v.firstname, v.lastname)
		for _, f := range v.ficfs {
			fmt.Println(f)
		}
	}
}
