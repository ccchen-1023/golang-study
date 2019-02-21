package main

import (
	"fmt"
)

func main() {

	cardSuits := [4]string{"Spades", "Hearts", "Diamonds"}
	// first argument to append must be slice; have [4]string
	// cardSuits = append(cardSuits, "Clubs")

	fmt.Println("for i, cardSuit := range cardSuits")
	for i, cardSuit := range cardSuits {
		fmt.Println(i, cardSuit)
	}

}
