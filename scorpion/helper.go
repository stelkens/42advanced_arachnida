package main

import (
	"path/filepath"
	"strings"
)

func isSupported(filename string) (bool) {
	
	ext := strings.ToLower(filepath.Ext(filename))

	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif":
		return true
	default:
		return false
	}
}