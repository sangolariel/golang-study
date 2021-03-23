package main

func main() {
	cards := deck{"Learning Go", returnString()}
	cards = append(cards, "Start try hard!")

	cards.print()
}

func returnString() string {
	return "Func return string type"
}
