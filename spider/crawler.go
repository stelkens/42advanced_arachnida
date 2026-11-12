package main

import (
	"net/http"
	"log"

	"golang.org/x/net/html"
)

//goes to the links on a webpage recursivly if max depth is not yet reached
//and if link was not yet visited
//adds all pictuer URL to map and saves a list of Links
//continue to crawle to the list of links
func crawl(client *http.Client, pageURL string, depth int, maxDepth int, visitedURLs map[string]bool, imagesURLs map[string]bool){
	
	if depth > maxDepth && len(visitedURLs) >= maxPagesVisited{
		return
	}

	if visitedURLs[pageURL]{
		return
	}

	visitedURLs[pageURL] = true


	//access Webpage
	response, err := client.Get(pageURL)
	
	if  err != nil {
		log.Println("Crawling failed:", err)
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
		crawl(client,link, depth + 1, maxDepth, visitedURLs, imagesURLs)
	}
}