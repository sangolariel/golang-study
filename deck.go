package main

import "fmt"

// Create a new type of Deck

type deck []string

func newDeck() deck {
	cards := deck{}
	cardsSuits := []string{"Heart", "Diamonds"}
	cardValue := []string{"One", "Two", "Three"}

	for _, suit := range cardsSuits {
		for _, val := range cardValue {
			cards = append(cards, val+"of"+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
