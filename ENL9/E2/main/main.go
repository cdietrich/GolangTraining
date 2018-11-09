package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func (p *Person) speak() {
	fmt.Println("Hello, my name is", p.Name)
}

// Human - super nice interface
type Human interface {
	speak()
}

func saySomething(h Human) {
	h.speak()
}

func main() {
	p := Person{"James Bond"}
	//saySomething(p)
	saySomething(&p)
}
