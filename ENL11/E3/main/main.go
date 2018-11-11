package main

import (
	"errors"
	"fmt"
)

type customErr struct {
	id      int
	message string
}

func (err customErr) Error() string {
	m := fmt.Sprintf("[%X] error: %v", err.id, err.message)
	return m
}

func main() {
	foo(customErr{123456789, "This is really bad"})
	fmt.Println("===============")
	foo(errors.New("A normal Error"))
}

func foo(err error) {
	fmt.Println(err)
	if v, ok := err.(customErr); ok {
		fmt.Println(v.id)
	}
}
