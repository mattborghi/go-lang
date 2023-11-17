package main

import (
	"fmt"
	"net/http"
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
	fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!!")
		// Send a value to the channel
		c <- "Might be down"
		return
	}

	fmt.Println(link, "is up!!")
	c <- "It's up"
}
