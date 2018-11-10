package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2)
	fmt.Printf("main: %T\n", c)
	go send(c)
	rec(c)
}

func send(c chan<- int) {
	fmt.Printf("send: %T\n", c)
	c <- 23
	c <- 42
}

func rec(c <-chan int) {
	fmt.Printf("rec: %T\n", c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
