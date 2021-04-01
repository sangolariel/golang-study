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
	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()
}

func (pointerPerson *person) updateName(newFirstName string) {
	(*pointerPerson).firstname = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
