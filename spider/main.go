package main

import (
	"os"
	"log"
	"net/http"
	"fmt"

	"golang.org/x/net/html"
)

func main() {


	// content, err := os.ReadFile("test.html")

	// if  err != nil {
	// 	log.Println("Accessing file failed", err)
	// 	return
	// }

	// //makes HTML-Tree

	// doc, err := html.Parse(strings.NewReader(string(content)))

	// if  err != nil {
	// 	log.Println("Converting HTML to Tree failed", err)
	// 	return
	// }

	if len(os.Args) <2{
		log.Println("Usage: go run . <URL>")
		return
	}

	url	:= os.Args[1]
	
	//access URL
	response, err := http.Get(url)
	
	if  err != nil {
		fmt.Println("Accessing URL failed")
		return
	}

	defer response.Body.Close()

	//check HTML Status
	if response.StatusCode != http.StatusOK {
		log.Println("HTML Error:", response.Status)
		return
	}

	//converts HTML-Text to DOM-Tree
	tree, err := html.Parse(response.Body)

	if  err != nil {
		log.Println("Converting HTML to Tree failed", err)
		return
	}
	count := 0;
	findImage(tree, url, &count)
}