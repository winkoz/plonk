package io

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/winkoz/plonk/internal/io/log"
)

type WalkFunc func(path string, info os.FileInfo, err error) error

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
	if !FileExists(path) {
		err := fmt.Errorf("File does not exist at path: %s", path)
		log.Error(err)
		return nil, err
	}

	data, err := ioutil.ReadFile(path)
	log.Error(string(data))
	if err != nil {
		log.Errorf("Error reading file %s: %+v", path, err)
		return []byte{}, err
	}
	return data, nil
}

// Walk walks the entire file structure for `root` and calls `walkFn` for each item it finds
func Walk(root string, walkFn WalkFunc) error {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		return walkFn(path, info, err)
	})

	if err != nil {
		log.Errorf("Error walking the path %q: %v\n", root, err)
		return err
	}

	return nil
}

// Append opens or creates file at `targetFilePath` and appends the `content` to it
func Append(targetFilePath string, content string) error {
	//Append second line
	file, err := os.OpenFile(targetFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	if err != nil {
		log.Errorf("Unable to open file %s. %v", targetFilePath, err)
		return err
	}
	if _, err := file.WriteString(content); err != nil {
		log.Errorf("Unable to append data to file %s. %v", targetFilePath, err)
		return err
	}

	return nil
}

func isValidPath(path string) error {
	_, err := os.Stat(path)
	return err
}
