package main

func main() {
	cards := newDeck()

	cards.shuffle()

	hand, remainingCards := deal(cards, 3)

	hand.print()

	remainingCards.print()
}
