package main

import "fmt"

func main() {
	if !true {
		fmt.Println("Guess what, this is NOT printed")
	}

	if !false {
		fmt.Println("Guess what, this won't NOT be printed")
	}

}
