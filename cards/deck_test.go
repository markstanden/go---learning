package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	deckSize := 52

	if len(d) != deckSize {
		t.Errorf("Expected deck length of %v, but got %v", deckSize, len(d))
	}
}

func TestFirstCard(t *testing.T) {
	d := newDeck()
	fc := "Ace of Spades"

	if d[0] != fc {
		t.Errorf("Expected first card of deck to be %s, but got %s", fc, d[0])
	}
}

func TestLastCard(t *testing.T) {
	d := newDeck()
	lc := "King of Diamonds"

	if d[len(d)-1] != lc {
		t.Errorf("Expected last card of deck to be %s, but got %s", lc, d[len(d)-1])
	}
}
