package main

import "fmt"

type pony int
type dog int

var myPony pony
var myDog dog

func main() {
	fmt.Printf("%v %T\n", myPony, myPony)
	myPony = 42
	fmt.Println(myPony)
	// myDog = myPony does not work
}
