package main

import (
	"fmt"
	"os"
)

func main() {
	defer df()
	_, err := os.Open("no-file.txt")
	if err != nil {
		panic(err)
	}
}

func df() {
	fmt.Println("i am deferred")
}
