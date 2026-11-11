package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync/atomic"
)

const maxPagesVisited = 100

func main() {

	//create new flag set after programmname
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	//definition of flags
	recursive := flag.Bool("r", false, "Recursive download")
	maxDepth := flag.Int("l", 5, "Maximum recursion depth")
	downloadPath := flag.String("p", "./data/", "Download Path")

	//gives flags there value
	flag.Parse()

	//sets args to argument, which are no flags
	args := flag.Args()

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

	fmt.Println("[start crawling]")

	//start crawling throug webside
	crawl(startURL, 0, *maxDepth, visitedURLs, images)

	//threat save counter
	var count atomic.Int64

	//create folder for the downloads
	err := os.MkdirAll(*downloadPath, 0755)
	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[start download]\n please wait ...")
	
	//start downloading pictures
	workerDownload(images, &count, *downloadPath)
	
	fmt.Println("[finish download]")
}