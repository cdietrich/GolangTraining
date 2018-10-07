package main

import "fmt"

var x int = 1

func increment() int {
	x++
	return x
}

func wrapper() func() int {
	x := 10
	return func() int {
		x++
		return x
	}
}

func main() {
	fmt.Println(increment())
	fmt.Println(increment())


	x := 0
	increment2 := func () int {
		x++
		return x
	}
	fmt.Println(increment2())
	fmt.Println(increment2())

	increment3 := wrapper()
	fmt.Println(increment3())
	fmt.Println(increment3())
}