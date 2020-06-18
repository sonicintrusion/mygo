package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// the first argument is the filename
	myfile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Danger Will Robinson! \nError:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, myfile)
}
