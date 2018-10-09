package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//a()
	b()
}

func a() {
	fmt.Println("xxxxx")
	result, error := http.Get("http://www.dietrich-it.de")
	if error != nil {
		log.Fatal(error)
	}
	page, error := ioutil.ReadAll(result.Body)
	result.Body.Close()
	if error != nil {
		log.Fatal(error)
	}
	fmt.Printf("%s", page)
}

func b() {
	fmt.Println("xxxxx")
	result, _ := http.Get("http://www.mcleods.com")
	page, _ := ioutil.ReadAll(result.Body)
	result.Body.Close()
	fmt.Printf("%s", page)
}
