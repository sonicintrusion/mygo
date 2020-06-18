package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func main() {
	//1st
	alex := person{"Alex", "Anderson"}
	paul := person{firstName: "Paul", lastName: "Baker"}

	fmt.Println(alex)
	fmt.Println(paul)

	//2nd
	var bob person

	fmt.Printf("%+v", bob)

	bob.firstName = "Bob"
	bob.lastName = "Johnson"

	fmt.Println(bob)

}
