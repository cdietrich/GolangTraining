package main

import (
	"fmt"
)

func main() {
	var i int
	f := func() {
		i++
	}
	f()
	f()
	f()
	fmt.Println(i)
	fmt.Println("============")
	var inc = incrementer()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())

}

func incrementer() func() int {
	var ix = 0
	return func() int {
		ix++
		return ix
	}
}
