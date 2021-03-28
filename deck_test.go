package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 9 {
		t.Errorf("Expected deck has lenght 9 but got %v", len(d))
	}
}
