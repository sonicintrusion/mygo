package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	fmt.Println("Begin")

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	fmt.Println("End")
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// something very specific for the english bot
	return "Hello"
}

func (spanishBot) getGreeting() string {
	// something very specific for spanish bot
	return "Hola"
}
