package main

import "fmt"

type deck []string

// so diff from OO

func newDeck() deck {
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	cardsOfDeck := []string{}
	for _, cardSuit := range cardSuits {
		for _, cardVal := range cardValues {
			card := cardVal + " of " + cardSuit
			cardsOfDeck = append(cardsOfDeck, card)
		}
	}

	return cardsOfDeck
}

// receiver function
func (d deck) printCards() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) deal(handSize int) (deck, deck) {
	handCards := d[0:handSize]
	// handCards := d[:handSize]
	remainCards := d[handSize:]
	return handCards, remainCards
}
