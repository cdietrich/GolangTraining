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

type human interface {
	speak()
}

func (p person) speak() {
	fmt.Println("Hello, i am", p.first, p.last)
}

func (p secretAgent) speak() {
	fmt.Println("Hello, i am", p.first, p.last, "with the ltk", p.ltk)
}

func letSpeak(h human) {
	h.speak()
}

func main() {
	p := person{first: "James", last: "Bond"}
	p.speak()
	letSpeak(p)
	var value human = human(p) // = p works too
	letSpeak(value)

	letSpeak(secretAgent{person: person{first: "X", last: "Y"}, ltk: true})

}

type dully struct {
}

// make sure we implement human for person
var _ human = (*person)(nil)

// compile error var _ human = (*dully)(nil)
