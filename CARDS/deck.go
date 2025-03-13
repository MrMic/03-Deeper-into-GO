package main

import (
	"fmt"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// ______________________________________________________________________
// Create a new deck of cards
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// ______________________________________________________________________
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// ______________________________________________________________________
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// ______________________________________________________________________
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
