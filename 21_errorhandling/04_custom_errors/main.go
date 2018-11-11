package main

import (
	"errors"
	"fmt"
)

// MyError is my superfancy error
type MyError struct {
	Message string
}

func (e *MyError) Error() string {
	return e.Message
}

// MyError2 is my superfancy error
type MyError2 struct {
	Message string
}

func (e MyError2) Error() string {
	return e.Message
}

func main() {
	var e1 = errors.New("This is bad")
	var e2 error = &MyError{"This is also bad"}
	var e3 error = MyError2{"This is terribly bad"}
	var e4 = fmt.Errorf("%v is bad, too", "This")
	fmt.Printf("%v\t%v\t%v\t%v\n", e1, e2, e3, e4)
	fmt.Printf("%T\t%T\t%T\t%T\n", e1, e2, e3, e4)
}
