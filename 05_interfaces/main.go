package main

import "fmt"

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
