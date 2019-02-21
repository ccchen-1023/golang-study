package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
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
	deckBytes := []byte(d.toString())
	return ioutil.WriteFile(filename, deckBytes, 0666)
}

func newDeckFromFile(filename string) deck {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		log.Fatal("Error: ", err)
	}
	return deck(strings.Split(string(bytes), ","))
}

func (d deck) shuffle() {
	epochNano := time.Now().UnixNano()
	source := rand.NewSource(epochNano)
	random := rand.New(source)

	randomSize := len(d) - 1

	for originalPosition := range d {
		// Intn returns, as an int, a non-negative pseudo-random number in [0,n) from the default Source. It panics if n <= 0.
		newPosition := random.Intn(randomSize)

		// swap
		originalVal := d[originalPosition]
		newValue := d[newPosition]
		d[originalPosition] = newValue
		d[newPosition] = originalVal

		// d[originalPosition], d[newPosition] = d[newPosition], d[originalPosition]

		// run at least 2 times and check if every output is same or diff?
	}
}
