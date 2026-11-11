package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
)

//creates a file with name <filename>
//copies download of <url> in that file
//returns error if it fails
func downloadImage(absoluteURL string, count *atomic.Int64, dowPath string) error{

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

	id:= count.Add(1)										//count thread safe

	filename := fmt.Sprintf("image_%d%s", id, imageFormat)	//creates filename
	
	file,err := os.Create(filepath.Join(dowPath,filename))

	if err != nil{
		log.Println("Error creating file:",err)
		return err
	}

	defer file.Close()

	_,err = io.Copy (file, response.Body)

	return err
}

func workerDownload(imagesURLs map[string]bool, count *atomic.Int64, dowPath string) {
	const workers = 4

	var wg sync.WaitGroup
	jobs := make(chan string)

	// register working workers
	for i := 0; i < workers; i++ {
		wg.Add(1)

		go func() {								//start worker
			defer wg.Done()						//tells if worker is done

			for imageURL := range jobs {		//waits for new URLs
				downloadImage(imageURL, count, dowPath)
				}
		}()
	}

	// gives URL to worker
	for imageURL := range imagesURLs {
		jobs <- imageURL
	}

	close(jobs)									//no more new URLs arriving
	wg.Wait()									//waits for all worker
}