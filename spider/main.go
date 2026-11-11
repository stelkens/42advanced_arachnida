package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

const maxPagesVisited = 100

func main() {

	//create new flag set after programmname
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	//definition of flags
	recursive := flag.Bool("r", false, "Recursive download")
	maxDepth := flag.Int("l", 5, "Maximum recursion depth")
	downloadPath := flag.String("p", "./data/", "Download Path")

	//if flag.Parse error exit on 1
	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
	os.Exit(1)
	}

	//sets args to argument, which are no flags
	args := flag.Args()

	//if more arguments than flags and URL exit on 1
	if len(args) != 1{
		log.Println("Usage: ./spider [-r] [-l N] [-p PATH] <URL>")
		os.Exit(1)
	}

	startURL	:= args[0]

	if !*recursive{
		*maxDepth = 0
	}

	//prepare maps
	visitedURLs := make(map[string]bool)
	images := make(map[string]bool)

	fmt.Println("[start crawling]\n please wait ...")

	//create client how waits max 10s for webside response
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	//start crawling throug webside
	crawl(client, startURL, 0, *maxDepth, visitedURLs, images)

	//threat save counter
	var count atomic.Int64

	//create folder for the downloads
	err := os.MkdirAll(*downloadPath, 0755)
	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[start download]\n please wait ...")
	
	//start downloading pictures
	workerDownload(client, images, &count, *downloadPath)
	
	fmt.Println("[finish download]")
}