package main

import (
	//"io"
	"log"
	"os"
)

//var out io.Writer =os.Stdout

func main(){

	if len(os.Args[1:]) < 1{
		log.Printf("Usage: ./scorpion FILE1 [FILE2 ...]")
		os.Exit(1)
	}
}