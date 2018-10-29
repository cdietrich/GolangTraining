package main

import (
	"fmt"
)

func main() {
	i := 0
	if i > 0 {
		fmt.Println("positive")
	} else if i == 0 {
		fmt.Println("zero")
	} else {
		fmt.Println("negative")
	}
}
