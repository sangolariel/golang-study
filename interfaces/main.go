package main

import "fmt"

type englishBot struct {
}

type spanisBot struct {
}

func main() {
	eb := englishBot{}
	sb := spanisBot{}

	eb.getGreeting()
	sb.getGreeting()
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb englishBot) {
	fmt.Println(sb.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi there"
}

func (spanisBot) getGreeting() string {
	return "Hola!"
}
