package main

import (
	"fmt"
)

func main() {

	colours := map[string]string{
		"red":   "#ff0000",
		"blue":  "#0000ff",
		"green": "#008000",
	}
	printMap(colours)

}

func printMap(c map[string]string) {
	for colour, hex := range c {
		fmt.Println("Hex code for", colour, "is", hex)
	}
}
