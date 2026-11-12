package main

import (
	//"io"
	"fmt"
	"log"
	"os"
)

//var out io.Writer =os.Stdout

func scorpion(path string, out *os.File) error{

	info, err := os.Stat(path)
	if err != nil{
		return err
	}

	if info.IsDir(){
		return fmt.Errorf("Error: %s is a directory", path)
	}

	_, err = fmt.Fprintf(out,
	"**File: %s**\nFilename: %s\nSize: %d Bytes\nLast Modified: %s\n\n",
	path,
	info.Name(),
	info.Size(),
	info.ModTime().Format("02.01.2006 15:04"),		//Format in Date and Time
	)

	if err != nil{
		return err
	}

	return nil
}

func main(){

	if len(os.Args[1:]) < 1{
		log.Printf("Usage: ./scorpion FILE1 [FILE2 ...]")
		os.Exit(1)
	}

	output, err := os.Create("output.txt")

	if err != nil{
		log.Println("Error creating output file:",err)
		os.Exit(1)
	}

	for _, path := range os.Args[1:]{
		
		if isSupported(path){
			
			err = scorpion(path, output)
			if err != nil{
				log.Printf("Extract data from %s failed: %s",path ,err)
			}
		}
	}

	defer output.Close()
}