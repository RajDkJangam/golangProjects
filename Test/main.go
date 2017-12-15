package main

import "fmt"

func main() {
	//printState()
	//greeting := "Hi there!!"

	//	go (func() {
	//	fmt.Println(greeting)
	//	})()
	c := make(chan string)
	c <- "Hi There"
	fmt.Println(<-c)
}
