package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Begin")

	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// io.Copy (writer, reader) interface
	io.Copy(os.Stdout, resp.Body) // Copy reads from resp and writes to Stdout

	fmt.Println("End")
}
