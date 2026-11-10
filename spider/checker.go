package main

import (
	"fmt"
	"log"
	"net/url"
	"path"
	"strings"
)

func valImage(imageURL string) (string, error){
	
	parsedURL,err := url.Parse(imageURL)

	if err != nil {
		log.Println("Invalid page URL:", err)
		return "",err
	}

	ext := strings.ToLower(path.Ext(parsedURL.Path))

	switch ext{
	case ".jpg", ".jpeg", ".png", ".gif", ".bmp":
		return ext,err
	default:
		return "", fmt.Errorf("Error: image format not valid")
	}
}