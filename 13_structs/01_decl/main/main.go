package main

import (
	"fmt"
)

type person struct {
	firstname string
	lastname  string
	age       int
}

func main() {
	p1 := person{
		firstname: "James",
		lastname:  "Bond",
		age:       32,
	}
	var p2 person = person{
		firstname: "Miss",
		lastname:  "Moneypenny",
		age:       27,
	}
	fmt.Println(p1, p2)
	fmt.Println(p1.firstname, p1.lastname, p1.age)
	fmt.Println(p2.firstname, p2.lastname, p2.age)
	p2.firstname = "Mrs"
	p2.lastname = "Bond"
	fmt.Println(p2.firstname, p2.lastname, p2.age)
}
