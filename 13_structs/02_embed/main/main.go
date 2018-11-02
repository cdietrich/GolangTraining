package main

import (
	"fmt"
)

type person struct {
	firstname string
	lastname  string
	age       int
}

type secretAgent struct {
	person
	ltk      bool
	codename string
}

type couple struct {
	person1 person
	person2 person
}

func main() {
	p1 := secretAgent{
		codename: "007",
		ltk:      true,
		person: person{
			firstname: "James",
			lastname:  "Bond",
			age:       32,
		},
	}
	var p2 person = person{
		firstname: "Miss",
		lastname:  "Moneypenny",
		age:       27,
	}
	fmt.Println(p1, p2)
	fmt.Println(p1.firstname, p1.lastname, p1.age, p1.ltk, p1.codename)
	fmt.Println(p2.firstname, p2.lastname, p2.age)
	p2.firstname = "Mrs"
	p2.lastname = "Bond"
	fmt.Println(p2.firstname, p2.lastname, p2.age)
	cpl := couple{
		person1: p1.person,
		person2: p2,
	}
	fmt.Println(cpl)
}
