package main

import (
	"fmt"
)

func main() {
	deck := newDeck()
	handCards, remainCards := deck.deal(1)
	fmt.Println("handCards: ")
	handCards.printCards()
	fmt.Println("remainCards: ")
	remainCards.printCards()
}
