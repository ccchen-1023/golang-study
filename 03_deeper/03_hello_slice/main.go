package main

import (
	"fmt"
)

func main() {

	cardSuits := []string{"Spades", "Hearts", "Diamonds"}
	cardSuits = append(cardSuits, "Clubs")

	fmt.Println("for i, cardSuit := range cardSuits")
	for i, cardSuit := range cardSuits {
		fmt.Println(i, cardSuit)
	}

	fmt.Println("for cardSuit := range cardSuits")
	for cardSuit := range cardSuits {
		fmt.Println(cardSuit)
	}

	fmt.Println("for _, cardSuit := range cardSuit")
	for _, cardSuit := range cardSuits {
		fmt.Println(cardSuit)
	}

}
