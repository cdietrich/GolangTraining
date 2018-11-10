package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	c := make(chan int)
	go foo(c)
	go bar(c)
	wg.Wait()
	fmt.Println("Done")
}

func foo(c chan<- int) {
	c <- 12
	wg.Done()
}

func bar(c <-chan int) {
	fmt.Println(<-c)
	wg.Done()
}
