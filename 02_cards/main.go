package main

func main() {
	filename := "./02_cards/deck.output"
	cards := newDeck()
	cards.saveToFile(filename)

	readCards := newDeckFromFile(filename)
	readCards.shuffle()
	readCards.print()
	// hand, remainingCards := deal(cards, 5)

	// fmt.Println("Hand: ")
	// hand.print()
	// fmt.Println("Remaining Cards: ")
	// remainingCards.print()
}
