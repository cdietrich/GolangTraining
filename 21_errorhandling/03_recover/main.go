package main

import (
	"errors"
	"fmt"
)

func main() {
	a()
	fmt.Println("after a")
}

func a() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in a:", r)
		}
	}()
	fmt.Println("Calling b.")
	b()
	fmt.Println("Returned normally from b.")

}

func b() {
	fmt.Println("am in b, gonna panic now")
	panic(errors.New("BOOM"))
}

func recoverMain() {
	fmt.Println("gonna revocer main")
	err := recover()
	if err != nil {
		fmt.Println("found error:", err)
	}
}
