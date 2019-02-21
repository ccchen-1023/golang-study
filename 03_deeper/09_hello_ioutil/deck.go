package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

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
	return deck(handCards), deck(remainCards)
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	// TODO: how to know the []byte built-in function?
	deckBytes := []byte(d.toString())
	return ioutil.WriteFile(filename, deckBytes, 0666)
}

func newDeckFromFile(filename string) deck {
	bytes, err := ioutil.ReadFile(filename)
	// error handlings
	if err != nil {
		fmt.Println("Error: ", err)
		log.Fatal("Error: ", err)
		// os.Exit(1)
		// TODO: where to find the code list?
		// os.Exit(9800001)

	}
	return deck(strings.Split(string(bytes), ","))
}
