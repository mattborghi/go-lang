package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingCards := deal(cards, 5)
	fmt.Println(hand.toString())

	fmt.Println("Hand: ")
	hand.print()
	fmt.Println("Remaining Cards: ")
	remainingCards.print()
}
