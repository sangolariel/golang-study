package main

import (
	"testing",
	"os"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 9 {
		t.Errorf("Expected deck has lenght 9 but got %v", len(d))
	}

	if d[0] !== "One of Heart" {
		t.Errorf("Expect first deck is One of Heart, but got %v", d[0])
	}

	if d[len(d) - 1] !== "Three of Diamonds" {
		t.Errorf("Expect last deck is Three of Diamonds, but got %v", d[len(d) -1])
	}
}

func TestSaveToDeckAndNewDeckFromFile (t *testing.T) {
	os.RemoveAll("_decktesting")

	deck := newDeck()

	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 9 {
		t.Errorf("Expect 9 card but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
