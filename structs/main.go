package main

import "fmt"

type contact struct {
	email   string
	zipCode int
}

type person struct {
	firstname string
	lastname  string
	contact
}

func main() {
	jim := person{
		firstname: "Jim",
		lastname:  "Party",
		contact: contact{
			email:   "jim@golang.io",
			zipCode: 700000,
		},
	}
	jim.updateName("Jimmy")
	jim.print()
}

func (p *person) updateName(newFirstName string) {
	p.firstname = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
