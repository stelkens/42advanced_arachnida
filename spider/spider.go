package main

import (
	"log"
	"fmt"
	"golang.org/x/net/html"
)

// goes recursevly through *html.Node; ++count for every found img
// finds an "img"-Element and the src of that "img"-Element
// calls func downloadImage with url = <src> and filename = <image_count.jpg>
func findImage(node *html.Node, pageURL string, count *int){
	if node.Type == html.ElementNode && node.Data == "img" {
		for _, attr := range node.Attr{
			if attr.Key == "src" {
				filename := fmt.Sprintf("image_%d.jpg", *count)
				err :=downloadImage(attr.Val, pageURL, filename)
				if err != nil {
					log.Println("Download failed", err)
				}else{
					*count++
				}
			}
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling{
		findImage(child, pageURL, count)
	}
}
