package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("mycards")
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()
	// deck := newDeckFromFile("mycards")
	// deck.print()
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
