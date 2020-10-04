package main

import "fmt"

func main() {

	cards := deck{"Ace of Diamond", "K f Diamond"}
	for i, card := range cards {
		fmt.Println(i, card)
	}

	// receiver
	//deck := newDeck()
	//deck.printCards()
	// go run main.go deck.go
}
