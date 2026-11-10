package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	//"net/url"
	"os"
)

//creates a file with name <filename>
//copies download of <url> in that file
//returns error if it fails
func downloadImage(absoluteURL string, count *int) error{

	response,err := http.Get(absoluteURL)					//access URL

	if err != nil {
		log.Println("Get response failed:", err)
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP error: %s", response.Status)
	}

	imageFormat := valImage(absoluteURL)						//checks if image has right format

	if imageFormat == ""{
		return nil
	}

	filename := fmt.Sprintf("image_%d%s", *count, imageFormat)	//creates filename
	
	file,err := os.Create(filename)

	if err != nil{
		log.Println("Error creating file:",err)
		return err
	}

	defer file.Close()

	_,err = io.Copy (file, response.Body)
	
	*count++

	return err
}