package main

// import format package
import (
	"fmt"
				"strings"
			)

// Constants
const SEPERATOR string = ","

// Create a new type 'deck'
//which is a slice of strings
type deck []string

/**
Creates a new deck
*/
func newDeck() deck {
	cards := deck {}
	cardSuits := []string {"Spades", "Hearts", "Clubs", "Diamonds"}
	cardValues  := []string {"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, (value + " of " + suit))
		}
	}
	return cards
}

/**
Deals a hand from the deck, returns the new hand, and the remainder of the deck
*/
func deal (d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

/**
Displays the content of the deck
*/
func (d deck) print () {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

/**
turns the deck into a single string, items separated by the SEPERATOR constant
*/
func (d deck) toString() string {
	return strings.Join([]string(d),SEPERATOR)
}