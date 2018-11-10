package main

import (
	"fmt"
)

func main() {
	// buffer
	c := make(chan int, 1)
	c <- 23
	fmt.Println(<-c)
	c = make(chan int, 2)
	c <- 23
	c <- 42
	fmt.Println(<-c)
	fmt.Println(<-c)
}
