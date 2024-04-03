package main

import (
	"fmt"
	"os"
)

func main() {
	// Path to file
	filePath := "first-post.txt"

	// Read file
	data, err := os.ReadFile(filePath)
	if err != nil {
		// A common use of `panic` is to abort if a function returns an error
		// value that we don’t know how to (or want to) handle. This example
		// panics if we get an unexpected error when creating a new file.
		panic(err)
	}

	// Print file contents
	fmt.Println(string(data))
}
