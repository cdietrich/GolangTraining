package main

import (
	"fmt"
)

type person struct {
	name string
}

func (p person) doit1(s string) {
	p.name = s
}

func (p *person) doit2(s string) {
	p.name = s
}

func main() {
	p := person{"Me"}
	fmt.Println(p)
	p.doit1("Too")
	fmt.Println(p)
	(&p).doit1("No Compile Error")
	fmt.Println(p)
	p.doit2("You")
	fmt.Println(p)
	(&p).doit2("Two")
	fmt.Println(p)
}
