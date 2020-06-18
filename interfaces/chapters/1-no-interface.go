package main

import (
	"fmt"
)

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

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb englishBot) {
	fmt.Println(sb.getGreeting())
}

func (englishBot) getGreeting() string {
	// something very specific for the english bot
	return "Hello"
}

func (spanishBot) getGreeting() string {
	// something very specific for spanish bot
	return "Hola"
}
