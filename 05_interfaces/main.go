package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	english := englishBot{}
	spanish := spanishBot{}

	printGreeting(english)
	printGreeting(spanish)

	resp, err := http.Get("http://gooogle.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// This is not the actual HTML body
	// fmt.Println("Response correct!", resp)
	// Instead of this we can use another approach
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println("Response correct!", string(bs))
	io.Copy(os.Stdout, resp.Body)
}

func (eb englishBot) getGreeting() string {
	// some custom and complex logic
	return "Hello there!"
}

func (sb spanishBot) getGreeting() string {
	// some custom and complex logic
	return "Hola ahi!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
