package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	jim.print()

	jim.updateName("jimmy")
	jim.print()

	fmt.Println("End")
}

func (p person) updateName(newFirstName string) {
	// this doesn't work -- something about pointers
	// Pass by value!!
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
