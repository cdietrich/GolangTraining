package main

import (
	"fmt"
)

func main() {
	var a1, a2, a3 string

	fmt.Print("Your Name: ")
	_, err := fmt.Scan(&a1)
	if err != nil {
		panic(err)
	}

	fmt.Print("Favourite Food: ")
	_, err = fmt.Scan(&a2)
	if err != nil {
		panic(err)
	}

	fmt.Print("Favourite Sport: ")
	_, err = fmt.Scan(&a3)
	if err != nil {
		panic(err)
	}
	fmt.Println(a1, a2, a3)
}
