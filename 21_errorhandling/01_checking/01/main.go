package main

import (
	"fmt"
)

func main() {
	n, err := fmt.Println("Hello error!")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n)
	}
}
