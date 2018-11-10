package main

func main() {
	// buffer
	c := make(chan<- int, 1)
	c <- 23
	// does not compile
	// fmt.Println(<-c)
}
