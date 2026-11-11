package main

import (
	"fmt"
	"net/http"
	"log"

	"golang.org/x/net/html"
)

func crawl(pageURL string, depth int, maxDepth int, visitedURLs map[string]bool, imagesURLs map[string]bool){
	
	if depth > maxDepth && len(visitedURLs) >= maxPagesVisited{
		return
	}

	if visitedURLs[pageURL]{
		return
	}

	visitedURLs[pageURL] = true


	//access Webpage
	response, err := http.Get(pageURL)
	
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

	var linksList []string

	//map all imageLinks
	findImages(tree, pageURL, imagesURLs)

	//list all links
	linksList = findLinks(tree, pageURL, linksList)

	for _,link := range linksList{
		crawl(link, depth + 1, maxDepth, visitedURLs, imagesURLs)
	}
}