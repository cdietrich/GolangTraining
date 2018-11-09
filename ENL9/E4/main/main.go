package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var mutex = &sync.Mutex{}
var value = 0

func inc() {
	mutex.Lock()
	copy := value
	copy++
	value = copy
	mutex.Unlock()
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
