package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstname string
	lastname  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstname: "Jim",
		lastname:  "Party",
		contact: contactInfo{
			email:   "jim@golang.io",
			zipCode: 700000,
		},
	}

	fmt.Printf("%+v", jim)
}
