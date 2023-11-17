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

	// Create a new channel to communicate go routines
	c := make(chan string)

	// Verify if each domain is respoding to traffic
	for _, link := range links {
		// The `go` tag makes the function run on a separate go routine
		go checkLink(link, c)
	}

	// Receive a value from the channel
	// Shorthand for not storing the channel message in a intermediate variable
	// This is a blocking call. After the channel receives a value the main routine exits.
	// fmt.Println(<-c)
	// Instead of concatenating one instruction for every expected channel to be received we can write a for loop
	// for i := 0; i < len(links); i++ {
	// Use an infinite loop instead
	for l := range c {
		// Create a function literal (anonymous functions or lambda)
		go func(link string) {
			time.Sleep(2 * time.Second)
			// rangign over c avoids writing <-c
			// fmt.Println(<-c)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	// We don't want a pause during the inital fetch
	// Wait two seconds to run this
	// time.Sleep(time.Second * 2)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!!")
		// Send a value to the channel
		// c <- "Might be down"
		// Send the link because we want to kick off another checkLink routine for that link
		c <- link
		return
	}

	fmt.Println(link, "is up!!")
	// c <- "It's up"
	c <- link
}
