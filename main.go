package main

func main() {
	cards := newDeck()

	hand, remainCard := deal(cards, 3)

	hand.print()
	remainCard.print()
}

func returnString() string {
	return "Func return string type"
}
