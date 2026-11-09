package main

import (
	"fmt"
	"golang.org/x/net/html"
)

// goes recursevly through *html.Node; ++count for every found img
// finds an "img"-Element and the src of that "img"-Element
// calls func downloadImage with url = <src> and filename = <image_count.jpg>
func findImage(node *html.Node, count *int){
	if node.Type == html.ElementNode && node.Data == "img" {
			for _, attr := range node.Attr{
				if attr.Key == "src" {
					downloadImage(attr.Val, fmt.Sprintf("%s%d.jpg", "image_", *count))
					*count++
				}
			}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling{
		findImage(child, count)
	}
}

