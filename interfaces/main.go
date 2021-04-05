package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {
}

type spanisBot struct {
}

func main() {
	eb := englishBot{}
	sb := spanisBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi there"
}

func (spanisBot) getGreeting() string {
	return "Hola!"
}
