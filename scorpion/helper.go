package main

import (
	"log"
	"path/filepath"
	"strings"
)

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