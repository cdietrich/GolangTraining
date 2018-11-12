package main

import (
	"fmt"

	"github.com/cdietrich/GolangTraining/ENL13/E2/quote"
	"github.com/cdietrich/GolangTraining/ENL13/E2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
