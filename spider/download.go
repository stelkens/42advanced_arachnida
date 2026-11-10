package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

//creates a file with name <filename>
//copies download of <url> in that file
//returns error if it fails
func downloadImage(imageURL string, pageURL string, filename string) error{

	baseURL,err := url.Parse(pageURL)

	if err != nil {
		log.Println("Invalid page URL:", err)
		return err
	}

	relativeURL,err := url.Parse(imageURL)

	if err != nil {
		log.Println("Invalid page URL:", err)
		return err
	}

	//creates absolute URL out of page URL and image URL
	absoluteURL := baseURL.ResolveReference(relativeURL)

	response,err := http.Get(absoluteURL.String())

	if err != nil {
		log.Println("Get response failed:", err)
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP error: %s", response.Status)
	}

	file,err := os.Create(filename)

	if err != nil{
		log.Println("Error creating file:",err)
		return err
	}

	defer file.Close()

	_,err = io.Copy (file, response.Body)

	return err
}