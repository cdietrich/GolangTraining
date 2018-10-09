package main

import "fmt"

func zero(x int) {
	fmt.Println("x in zero", &x)
	x = 0
}

func main() {
	x := 5
	fmt.Println("x in main", &x)
	zero(x)
	fmt.Println(x) // = 5
}
