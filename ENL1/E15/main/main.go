package main

import "fmt"

type x int

var y x

func main() {
	fmt.Printf("%v %T\n", y, y)
	y = 42
	fmt.Println(y)
	z := int(y)
	fmt.Printf("%v %T\n", z, z)
}
