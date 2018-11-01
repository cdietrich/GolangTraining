package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	delete(m, `no_dr`)
	for person, things := range m {
		fmt.Println("Person:", person)
		for i, v := range things {
			fmt.Println(i, v)
		}
	}
}
