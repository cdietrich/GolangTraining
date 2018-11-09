package main

import (
	"fmt"
)

type thing struct {
}

func (t thing) doit() {
	fmt.Println("doit")
}

func (t *thing) doitptr() {
	fmt.Println("doit ptr")
}

type doer interface {
	doit()
}

type doerptr interface {
	doitptr()
}

func main() {
	var a thing = thing{}
	a.doit()
	a.doitptr()
	var b *thing = &a
	b.doit()
	b.doitptr()

	var x1 doer = a
	var x2 doer = b
	// does not work, method only defined for ptr
	//var x3 doerptr = a
	var x4 doerptr = b
	fmt.Println(x1, x2, x4)
}
