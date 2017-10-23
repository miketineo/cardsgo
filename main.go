package main

func main() {
	cards := newDeck()
	cards.saveToFile("mycards")
	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()
}
