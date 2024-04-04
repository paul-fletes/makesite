package main

import (
	"fmt"
	"os"
	"text/template"
)

// Page holds all the information needed to generate a new HTML page
// from a text file on the filesystem
type Page struct {
	TextFilePath string
	TextFileName string
	HTMLPagePath string
	Content      string
}

func main() {
	// Path to file
	filePath := "first-post.txt"

	// Read file
	content, err := os.ReadFile(filePath)
	if err != nil {
		// A common use of `panic` is to abort if a function returns an error
		// value that we don’t know how to (or want to) handle. This example
		// panics if we get an unexpected error when creating a new file.
		panic(err)
	}

	// Create a page
	page := Page{
		TextFilePath: "first-post.txt",
		TextFileName: "first-post",
		HTMLPagePath: "first-post.html",
		Content:      string(content),
	}

	// Create a new template in memory
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	// Create a blank HTML file
	newFile, err := os.Create("first-post.html")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	// Save the template insted created file
	err = t.Execute(newFile, page)
	if err != nil {
		panic(err)
	}

	fmt.Println("HTML template written to first-post.html")
}
