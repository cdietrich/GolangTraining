package main

import (
	"fmt"
)

// does not run
func main() {
	c := make(chan int)
	c <- 1
	fmt.Println(<-c)
}
