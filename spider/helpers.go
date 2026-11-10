package main

import (
	"net/url"
	"log"
	"mime"
	"net/http"
	"strings"
)

//checks if image has .jpg, .jpeg, .png, .gif or .bmp ending
//returns the format
// func valImage(imageURL string) (string){

// 	parsedURL,err := url.Parse(imageURL)

// 	if err != nil {
// 		log.Println("Invalid page URL:", err)
// 		return ""
// 	}

// 	ext := strings.ToLower(path.Ext(parsedURL.Path))

// 	switch ext{
// 	case ".jpg", ".jpeg", ".png", ".gif", ".bmp":
// 		return ext
// 	default:
// 		return ""
// 	}
// }

//checks if content type from img URL is valid
//returns image type as ".<type>" string 
func valImage(imageURL string) (string){
	
	response,err := http.Get(imageURL)

	if err != nil {
		log.Println("Accessing URL failed:", err)
		return ""
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return ""
	}

	contentType := response.Header.Get("Content-Type")

	mediaType,_,err := mime.ParseMediaType(contentType)

	if err != nil {
		log.Println("Content-Type error:", err)
		return ""
	}

	switch mediaType{
	case "image/jpg", "image/jpeg", "image/png", "image/gif", "image/bmp":
		ret := strings.TrimPrefix(mediaType, "image/")
		ret = "." + ret
		return ret
	default:
		return ""
	}
}

//creates absolute URL out of page URL and image URL
func getAbsolutPath(imageURL string, pageURL string) (string){
	baseURL,err := url.Parse(pageURL)

	if err != nil {
		log.Println("Invalid page URL:", err)
		return ""
	}

	relativeURL,err := url.Parse(imageURL)

	if err != nil {
		log.Println("Invalid page URL:", err)
		return ""
	}

	absoluteURL := baseURL.ResolveReference(relativeURL)

	return absoluteURL.String()
}