package main

import "fmt"

func main() {
	// buffer
	c := make(<-chan int, 1)
	// does not compile
	// c <- 23
	fmt.Println(<-c)
}
