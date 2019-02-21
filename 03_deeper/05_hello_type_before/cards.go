package main

import "fmt"

func newDeck() []string {
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

func printCards(cards []string) {
	for i, card := range cards {
		fmt.Println(i, card)
	}
}
