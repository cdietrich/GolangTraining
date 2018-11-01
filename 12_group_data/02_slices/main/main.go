package main

import (
	"fmt"
)

func main() {
	// decl
	a := []int{1, 2, 3, 4, 5}
	// access
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(a[1])
	for i, ai := range a {
		fmt.Println(ai, "at index", i)
	}
	// slicing
	b := []int{5, 6, 7, 8, 9}
	fmt.Println(b)
	fmt.Println(b[:])
	fmt.Println(b[1:3])
	fmt.Println(b[1:])
	fmt.Println(b[:3])
	// append and delete
	c := append(b, 23, 42)
	fmt.Println(b)
	fmt.Println(c)
	// unfold
	d := append([]int{1, 2}, []int{3, 4}...)
	fmt.Println(d)
	// delete**
	e := []int{1, 2, 3, 4, 5}
	fmt.Println(e)
	f := append(e[:2], e[3:]...)
	fmt.Println(f)
	// make
	g := make([]int, 0, 10)
	fmt.Println(g)
	fmt.Println(len(g))
	fmt.Println(cap(g))
	g = append(g, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(g)
	fmt.Println(len(g))
	fmt.Println(cap(g))
	h := make([]int, 0, 1)
	i := append(h, 1)
	j := append(h, 2)
	fmt.Print(h, len(h), cap(h))
	fmt.Print(i, len(i), cap(i))
	fmt.Print(j, len(j), cap(j))
	// multi-dim slices
	s1 := []string{"a", "b", "c"}
	s2 := []string{"x", "y", "z"}
	ss := [][]string{s1, s2}
	fmt.Println(ss)
	ss2 := [][]string{{"Hannibal", "Faceman", "BA", "Murdoch"}, {"Michael Knight"}}
	fmt.Println(ss2)
	// oob fmt.Println(ss2[1][2])
}
