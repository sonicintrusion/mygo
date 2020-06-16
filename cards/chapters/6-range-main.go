package main

import "fmt"

func main() {
	cards := newDeck()

	hand, reminingCards := deal(cards, 5)

	fmt.Println("My Hand")
	hand.print()

	fmt.Println("Remaining Deck")
	reminingCards.print()
}
