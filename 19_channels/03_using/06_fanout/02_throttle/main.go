package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go produceTasks(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func produceTasks(taskQueue chan<- int) {
	for i := 0; i < 100; i++ {
		taskQueue <- i
	}
	close(taskQueue)
}

func fanOutIn(taskQueue <-chan int, resultQueue chan<- int) {
	var wg sync.WaitGroup
	const goroutines = 10
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			for taskNumber := range taskQueue {
				resultQueue <- timeConsumingWork(taskNumber)
			}
			wg.Done()
		}()
	}

	for task := range taskQueue {
		wg.Add(1)
		go func(taskNumber int) {
			resultQueue <- timeConsumingWork(taskNumber)
			wg.Done()
		}(task)
	}
	wg.Wait()
	close(resultQueue)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
