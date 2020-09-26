package io

import (
	"io/ioutil"
	"os"

	"github.com/winkoz/plonk/internal/io/log"
)

// GetCurrentDir returns the directory in which the project is running.
func GetCurrentDir() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return path
}

// DirectoryExists returns wether or not a file is found on disk
func DirectoryExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// FileExists returns wether or not a file is found on disk
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// CreatePath in the local filesystem
func CreatePath(path string) error {
	err := os.MkdirAll(path, 0755)
	if !os.IsExist(err) {
		log.Errorf("CreatePath %s failed. %v", path, err)
		return err
	}
	return nil
}

// DeletePath in the local filesystem
func DeletePath(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		panic(err)
	}
}

// ReadFile reads a file
func ReadFile(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	log.Error(string(data))
	if err != nil {
		log.Errorf("Error reading file %s: %+v", path, err)
		return []byte{}, err
	}
	return data, nil
}

func isValidPath(path string) error {
	_, err := os.Stat(path)
	return err
}
