package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("name: %v age:%v", p.name, p.age)
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] } // tuple assignment
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }

func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

	xp := []Person{Person{"Max", 72}, Person{"Lea", 10}, Person{"Udo", 42}}
	fmt.Println(xp)
	sort.Sort(ByAge(xp))
	fmt.Println(xp)

	xp = []Person{Person{"Max", 72}, Person{"Lea", 10}, Person{"Udo", 42}}
	fmt.Println(xp)
	sort.Slice(xp, func(i, j int) bool {
		return xp[i].name < xp[j].name
	})
	fmt.Println(xp)
	fmt.Println("a" < "b")
}
