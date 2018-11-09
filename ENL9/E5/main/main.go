package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

var value int64 = 0

func inc() {
	atomic.AddInt64(&value, 1)
	wg.Done()
}

// go run -race main.go
func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go inc()
	}

	wg.Wait()
	fmt.Println(value)
}
