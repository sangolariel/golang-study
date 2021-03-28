package main

func main() {
	cards := newDeckFromFile("my_cards")
	cards.shuffe()
	cards.print()
}

func returnString() string {
	return "Func return string type"
}
