package main

import (
	"fmt"
)

func main() {
	fmt.Println("Begin")

	colours := map[string]string{
		"red":   "#ff0000",
		"blue":  "#0000ff",
		"green": "#008000",
	}

	fmt.Println(colours)

	printMap(colours)

	fmt.Println("End")
}

func printMap(c map[string]string) {
	for colour, hex := range c {
		fmt.Println("Hex code for", colour, "is", hex)
	}
}
