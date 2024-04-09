package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
	"strings"
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
	// Define a 'file' flag to specify input text file name
	fileName := flag.String("file", "first-post.txt", "Name of the input .txt file")
	flag.Parse()

	// Trim file extension from file name
	textFileName := strings.TrimSuffix(*fileName, ".txt")

	// Create a page
	page := Page{
		TextFilePath: *fileName,
		TextFileName: textFileName,
		HTMLPagePath: textFileName + ".html",
		Content:      "",
	}

	// Read in contents of input text file
	fileContents, err := os.ReadFile(page.TextFilePath)
	if err != nil {
		panic(err)
	}

	// Pass file contents to the page instance
	page.Content = string(fileContents)

	// Create a new template in memory
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	// Create an appropriately named HTML file
	htmlFile, err := os.Create(page.HTMLPagePath)
	if err != nil {
		panic(err)
	}
	defer htmlFile.Close()

	// Save the template inside created file
	err = t.Execute(htmlFile, page)
	if err != nil {
		panic(err)
	}

	fmt.Printf("HTML template written to %s\n", page.HTMLPagePath)
}
