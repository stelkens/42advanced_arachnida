package main

import (
	"log"
	"golang.org/x/net/html"
	"slices"
)

// goes recursevly through *html.Node; ++count for every found img
// finds an "img"-Element and the src of that "img"-Element
// calls func downloadImage with url = <src> and filename = <image_count.jpg>
func findImage(node *html.Node, pageURL string, count *int){
	if node.Type == html.ElementNode && node.Data == "img" {
		for _, attr := range node.Attr{
			if attr.Key == "src" {
				absolutPath := getAbsolutPath(attr.Val, pageURL)
				// imageFormat := valImage(absolutPath)				//checks if image has right format
				// if imageFormat == ""{
				// 	continue
				// }
				// filename := fmt.Sprintf("image_%d%s", *count, imageFormat)
				err := downloadImage(absolutPath, count)
				if err != nil {
					log.Println("Download failed", err)
				}
			}
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling{
		findImage(child, pageURL, count)
	}
}

//saves links recuvsivly in a []string without duplicates 
func findLinks(node *html.Node, pageURL string, listURL []string)([]string){
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr{
			if attr.Key == "href" {
				absolutPath := getAbsolutPath(attr.Val, pageURL)
				exists := slices.Contains(listURL, absolutPath)		//checkt if absolutPath in list is
				if !exists{
					listURL = append(listURL, absolutPath)
				}
			}
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling{
		listURL = findLinks(child, pageURL, listURL)
	}
	return listURL
}