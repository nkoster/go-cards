package main

import "testing"

const CARDSNEEDED = 52
const FIRSTCARD = "Ace of Hearts"
const LASTCARD = "King of Clubs"

func TestNewDeck(t *testing.T) {

	d := newDeck()
	if len(d) != CARDSNEEDED {
		t.Errorf("\nDeck does not hold %d cards but %v", CARDSNEEDED, len(d))
	}

	if d[0] != FIRSTCARD {
		t.Errorf("\nDeck does not have %s as the first card, but %v", FIRSTCARD, d[0])
	}

	if d[len(d)-1] != LASTCARD {
		t.Errorf("\nDeck does not have %s as the last card, but %v", LASTCARD, d[len(d)-1])
	}

}
