package main

import (
	"fmt"
)

func main() {
	v1 := 1 == 2
	v2 := 1 <= 2
	v3 := 1 >= 2
	v4 := 1 != 2
	v5 := 1 < 2
	v6 := 1 > 2
	fmt.Println(v1, v2, v3, v4, v5, v6)

}
