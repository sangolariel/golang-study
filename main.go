package main

import "fmt"

func main() {
	// cards := newDeck()

	// hand, remainCard := deal(cards, 3)

	// hand.print()
	// remainCard.print()
	// gretting := "Learning Go"

	// fmt.Println([]byte(gretting))

	cards := newDeck()

	cards.saveToFile("my_cards")

	fmt.Println(cards.toString())
}

func returnString() string {
	return "Func return string type"
}
