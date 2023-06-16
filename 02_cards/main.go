package main

import "fmt"

func main() {
	cards := newDeck()
	cards.saveToFile("./02_cards/deck.output")

	hand, remainingCards := deal(cards, 5)

	fmt.Println("Hand: ")
	hand.print()
	fmt.Println("Remaining Cards: ")
	remainingCards.print()
}
