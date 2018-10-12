package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i%3 == 0 {
			continue
		}
		fmt.Println(i)
		if i >= 22 {
			break
		}
	}
}
