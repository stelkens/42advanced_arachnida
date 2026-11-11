package main

import (
	"golang.org/x/net/html"
	"slices"
)

// goes recursevly through *html.Node; ++count for every found img
// finds an "img"-Element and the src of that "img"-Element
// calls func downloadImage with url = <src> and filename = <image_count.jpg>
func findImages(node *html.Node, pageURL string, images map[string]bool){
	if node.Type == html.ElementNode && node.Data == "img" {
		for _, attr := range node.Attr{
			if attr.Key == "src" {
				absolutePath, _ := getAbsolutPath(attr.Val, pageURL)

				if !images[absolutePath] {							// writes new URL in map
					images[absolutePath] = true
				}
			}
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling{
		findImages(child, pageURL, images)
	}
}

//saves links recursivly in a []string without duplicates 
func findLinks(node *html.Node, pageURL string, listURL []string)([]string){
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr{
			if attr.Key == "href" {
				absolutePath, sameDoman := getAbsolutPath(attr.Val, pageURL)
				exists := slices.Contains(listURL, absolutePath)		//checkt if absolutPath in list is
				if !exists && sameDoman{
					listURL = append(listURL, absolutePath)
				}
			}
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling{
		listURL = findLinks(child, pageURL, listURL)
	}
	return listURL
}