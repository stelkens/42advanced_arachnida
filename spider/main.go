package main

import (
	"os"
	"log"
	"golang.org/x/net/html"
	"strings"
)

func main() {


	content, err := os.ReadFile("test.html")

	if  err != nil {
		log.Println("Accessing file failed", err)
		return
	}

	//makes HTML-Tree

	doc, err := html.Parse(strings.NewReader(string(content)))

	if  err != nil {
		log.Println("Converting HTML to Tree failed", err)
		return
	}

	count := new(int)

	findImage(doc, count)
}