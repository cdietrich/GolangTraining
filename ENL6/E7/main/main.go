package main

import (
	"fmt"
)

func main() {
	a := sayHello
	a()
	b := func() {
		fmt.Println("Hello Worl 2")
	}
	b()
}

func sayHello() {
	fmt.Println("Hello World")
}
