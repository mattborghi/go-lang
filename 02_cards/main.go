package main

import "fmt"

func main() {
	filename := "./02_cards/deck.output"
	cards := newDeck()
	cards.saveToFile(filename)

	readCards := newDeckFromFile(filename)
	readCards.shuffle()

	hand, remainingCards := deal(readCards, 5)

	fmt.Println("Hand: ")
	hand.print()
	fmt.Println("Remaining Cards: ")
	remainingCards.print()
}
