package main

import (
	"fmt"
)

func main() {
	var favSport string = "hiking"
	switch favSport {
	case "running":
		fmt.Println("Meh")
	case "dancing":
		fmt.Println("Meh")
	case "hiking":
		fmt.Println("Yeah")
	default:
		fmt.Println("Ups")
	}
}
