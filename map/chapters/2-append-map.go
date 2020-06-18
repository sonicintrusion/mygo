package main

import (
	"fmt"
)

func main() {
	fmt.Println("Begin")

	// var colours map[string]string

	colours := make(map[string]string)

	// colours := map[string]string{
	// 	"red":   "#ff0000",
	// 	"blue":  "#0000ff",
	// 	"green": "#008000",
	// }

	colours["white"] = "#ffffff"
	fmt.Println(colours)

	delete(colours, "white")
	fmt.Println(colours)

	fmt.Println("End")
}
