package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(producer("Joe"), producer("Ann"))
	for i := 0; i < 100; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func producer(msg string) <-chan string {
	c := make(chan string)
	// will produce infite numer of messages to the channel
	// and then sleep randomly
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(i1, i2 <-chan string) <-chan string {
	c := make(chan string)
	// worker will read from given i and write on c
	worker := func(i <-chan string) {
		for {
			value := <-i
			c <- value
		}
	}
	go worker(i1)
	go worker(i2)
	return c
}
