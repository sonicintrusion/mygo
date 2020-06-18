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

	fmt.Println("End")
}
