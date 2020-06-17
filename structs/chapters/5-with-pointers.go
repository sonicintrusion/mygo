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

	jimPointer := &jim // & means "memory address of this variable"
	jimPointer.updateName("Jimmy")
	jim.print()

	fmt.Println("End")
}

// *person here represents a "pointer to a person"
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName // *pointer get the value inside the address (of the pointer)
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
