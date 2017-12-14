package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 { // In case file name is not provided
		fmt.Println("Please provide the file name.")
		os.Exit(1)
	}
	file, err := os.Open(os.Args[1]) // For read access.
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}
