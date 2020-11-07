package main

import "testing"

func TestNewDeck(t *testing.T) {
	const CardsNeeded = 52
	d := newDeck()
	if len(d) != CardsNeeded {
		t.Errorf("Deck does not hold %d cards but %v", CardsNeeded, len(d))
	}
}
