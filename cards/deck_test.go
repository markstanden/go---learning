package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	deckSize := 52

	if len(d) != deckSize {
		t.Errorf("Expected deck length of %v, but got %v", deckSize, len(d))
	}
	fc := "Ace of Spades"

	if d[0] != fc {
		t.Errorf("Expected first card of deck to be %s, but got %s", fc, d[0])
	}
	lc := "King of Diamonds"

	if d[len(d)-1] != lc {
		t.Errorf("Expected last card of deck to be %s, but got %s", lc, d[len(d)-1])
	}
}

func TestSaveToFileAndReadFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := readFromFile("_decktesting")

	if len(loadedDeck) != len(deck) {
		t.Errorf("Expected %v cards in deck, got %v", len(deck), len(loadedDeck))
	}

	os.Remove("_decktesting")
}
