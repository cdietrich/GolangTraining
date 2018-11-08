package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

type lower struct {
	value string
	Value string
}

type Thing struct {
	Name string
}

type ThingWithTag struct {
	Name string `json:"TotallyDifferent"`
}

type MyWriter struct {
}

func (w MyWriter) Write(p []byte) (n int, err error) {
	return os.Stdout.Write(p)
}

type MyReader struct {
	delegate io.Reader
}

func (r MyReader) Read(p []byte) (n int, err error) {
	nx, errx := r.delegate.Read(p)
	return nx, errx
}

func main() {
	v, _ := json.Marshal(1)
	os.Stdout.Write(v)
	p := Person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	v, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Stdout.Write(v)
	fmt.Printf("%T\n", v)
	fmt.Println(string(v))

	var pc Person
	err = json.Unmarshal(v, &pc)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pc)
	}
	var pc2 Person
	err = json.Unmarshal([]byte(`{"First":"James","Last":"Bond","Age":32}`), &pc2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pc2)
	}

	var pc3 Person
	err = json.Unmarshal([]byte(`{"Error","Last":"Bond","Age":32}`), &pc3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pc3)
	}

	// lowercase no export
	v, err = json.Marshal(lower{"demo", "demo2"})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(v))
	}

	// slice

	v, err = json.Marshal([]Thing{Thing{"X"}, Thing{"Y"}})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(v))
	}

	// map
	m := map[int]string{1: "x", 2: "y"}
	v, err = json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(v))
		var u map[int]string
		err = json.Unmarshal(v, &u)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(m, u)
		}
	}
	twt := ThingWithTag{"Demo"}
	v, err = json.Marshal(twt)
	fmt.Println(string(v))

	// encode to writer
	err = json.NewEncoder(os.Stdout).Encode(twt)
	if err != nil {
		fmt.Println(err)
	}

	err = json.NewEncoder(MyWriter{}).Encode(twt)
	if err != nil {
		fmt.Println(err)
	}

	var reader io.Reader = MyReader{strings.NewReader(`{"First":"James","Last":"Bond","Age":32}`)}
	dec := json.NewDecoder(reader)
	for {
		var px Person
		if err := dec.Decode(&px); err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s: %s\n", px.First, px.Last)
	}

	fmt.Fprintln(os.Stdout, "Hello World")
}
