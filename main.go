package main

import "fmt"

func main() {
	cards := []string{"Learning Go", returnString()}
	cards = append(cards, "Start try hard!")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func returnString() string {
	return "Func return string type"
}
