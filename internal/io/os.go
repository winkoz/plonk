package io

import (
	"log"
	"os"
)

// GetCurrentDir returns the directory in which the project is running.
func GetCurrentDir() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return path
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func isValidPath(path string) error {
	_, err := os.Stat(path)
	return err
}
