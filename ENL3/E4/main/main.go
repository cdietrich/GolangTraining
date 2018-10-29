package main

import (
	"fmt"
)

func main() {
	i := 1982
	for {
		fmt.Println(i)
		i++
		if i > 2018 {
			break
		}
	}
}
