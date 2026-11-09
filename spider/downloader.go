package main

import (
	"os"
	"log"
	"net/http"
	"io"
	"fmt"
)

//creates a file with name <filename>
//copies download of <url> in that file
//returns error if it fails
func downloadImage(url string, filename string) error{
	
	response,err := http.Get(url)

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