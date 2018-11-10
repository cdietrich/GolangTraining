package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fanIn := make(chan int)
	producers := []<-chan int{
		producer(0),
		producer(1),
		producer(2),
		producer(3),
		producer(4),
		producer(5),
		producer(6),
		producer(7),
		producer(8),
		producer(9),
	}
	fmt.Println("number of producers", len(producers))
	go receiver(producers, fanIn)
	for v := range fanIn {
		fmt.Println(v)
	}
}

func receiver(producers []<-chan int, fanIn chan<- int) {
	var wg sync.WaitGroup
	wg.Add(len(producers))
	for _, p := range producers {
		go func(px <-chan int) {
			for v := range px {
				fanIn <- v
				time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			}
			wg.Done()
		}(p)
	}
	wg.Wait()
	close(fanIn)
}

func producer(base int) chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- base*10 + i
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}
