package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}
type myslice []int
type bot interface { // all the types of the program
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
	var slice myslice
	printGreeting(slice)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

/*func printGreeting(sb spanishBot){
  fmt.Println(sb.getGreeting())
}*/

func (englishBot) getGreeting() string {
	//
	return "Hi There!!"
}

func (spanishBot) getGreeting() string {
	//
	return "Hola!!"
}

func (myslice) getGreeting() string {
	//
	return "Seriously!!"
}
