package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checklink(link, c)
	}
	/*for { // Stay alive request so no end points
		//fmt.Println(<-c)
		go checklink(<-c, c) //blocking syntax
	}*/
	for l := range c { // code refractor for the above code
		//fmt.Println(<-c)
		//time.Sleep(5*time.Second) //Not here as it will pause the Main routine
		// it may happen that other child routine is completed and adds value to channel.
		//go checklink(l, c)
		go func(link string) {
			time.Sleep(5 * time.Second)
			checklink(link, c)
		}(l)
	}
}

func checklink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		//c <- "Might be down I think"
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	//c <- "Yep its up"
	c <- link
}
