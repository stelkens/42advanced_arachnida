package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	_ "golang.org/x/image/bmp"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
)

var out io.Writer =os.Stdout

type printer struct {}

//Methode: reads exif data and makes readable design for the output
func (printer) Walk(name exif.FieldName, tag *tiff.Tag) error {
		_,err := fmt.Fprintf(out, " %s: %s\n", name, tag.String())
		return err
}

//checks if file ending is supported
func isSupported(filename string) (bool) {
	
	ext := strings.ToLower(filepath.Ext(filename))

	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif":
		return true
	default:
		log.Println("File type not supported:", filename)
		return false
	}
}

func scorpion(path string, out *os.File) error{

	info, err := os.Stat(path)
	if err != nil{
		return err
	}

	if info.IsDir(){
		return fmt.Errorf("Error: %s is a directory", path)
	}

	_, err = fmt.Fprintf(out,
	"\t***\tFile: %s\t***\nFilename: %s\nSize: %d Bytes\nLast Modified: %s\n",
	path,
	info.Name(),
	info.Size(),
	info.ModTime().Format("02.01.2006 15:04"),		//Format in Date and Time
	)

	if err != nil{
		return err
	}

	file, err := os.Open(path)

	if err != nil{
		return err
	}

	//Reads first Bytes of file and get some Metadata like format
	config, format, err := image.DecodeConfig(file)

	if err !=nil {
		return err
	}

	_, err = fmt.Fprintf(out,
	"Width: %d\nHeight: %d\nFormat: %s\n\n",
	config.Width,
	config.Height,
	format,
	)

	if err != nil{
		return err
	}

	file.Seek(0,io.SeekStart)				//reset reading position

	mData, err := exif.Decode(file)			//reads EXIF data

	if err != nil{
		fmt.Fprintln(out, " No EXIF data \n")
		return nil
	}

	//print creation time if it exists
	if time, err := mData.DateTime(); err == nil {
		fmt.Fprintf(out, "Creation date: %s\n",time.Format("02.01.2006 15:04"))
	}

	err = mData.Walk(printer{})

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

	// output is global variable out 
	out = output

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
			fmt.Fprintln(out, "\t\t__________________\n\n")
		}
	}

	defer output.Close()
}
