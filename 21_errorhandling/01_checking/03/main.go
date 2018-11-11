package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	io.Copy(f, strings.NewReader("James Bond, Miss Moneypenny"))

}
