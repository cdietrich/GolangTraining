package main

import (
	"fmt"
)

const (
	Y2014 = iota + 2014
	Y2015 = iota + 2014
	Y2016 = iota + 2014
	Y2017 = iota + 2014
	Y2018 = iota + 2014
)

func main() {
	fmt.Println(Y2014, Y2015, Y2016, Y2017, Y2018)
}
