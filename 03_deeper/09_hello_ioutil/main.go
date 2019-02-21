package main

func main() {
	deck := newDeck()
	deck.saveToFile("my_deck")
	// newDeckFromFile("gg")
	deckFromFile := newDeckFromFile("my_deck")
	deckFromFile.printCards()
}
