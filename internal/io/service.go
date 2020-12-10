package io

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/winkoz/plonk/data"
	"github.com/winkoz/plonk/internal/io/log"
)

// WalkFunc (s Service)is the callback method for the Walk function
type WalkFunc func(path string, info os.FileInfo, err error) error

// Service encasulates io related methods to the app
type Service interface {
	GetCurrentDir() string
	DirectoryExists(filename string) bool
	FileExists(filename string) bool
	CreatePath(path string) error
	DeletePath(path string)
	ReadFile(path string) ([]byte, error)
	Walk(root string, walkFn WalkFunc) error
	Append(targetFilePath string, content string) error
	Write(targetFilePath string, content string) error
	IsValidPath(path string) error
	Base64Encode(v []byte) (string, error)
}

type service struct{}

// NewService creates a new Service object
func NewService() Service {
	return service{}
}

// GetCurrentDir returns the directory in which the project is running.
func (s service) GetCurrentDir() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return path
}

// DirectoryExists returns wether or not a file is found on disk
func (s service) DirectoryExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// FileExists returns wether or not a file is found on disk
func (s service) FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// CreatePath in the local filesystem
func (s service) CreatePath(path string) error {
	err := os.MkdirAll(path, OwnerPermission)
	if err != nil && !os.IsExist(err) {
		log.Errorf("CreatePath %s failed. %v", path, err)
		return err
	}
	return nil
}

// DeletePath in the local filesystem
func (s service) DeletePath(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		panic(err)
	}
}

// ReadFile reads a file
func (s service) ReadFile(path string) ([]byte, error) {
	var resData []byte
	var err error

	if strings.Contains(path, BinaryFile) {
		binaryPath := strings.TrimPrefix(path, BinaryFile+"/")
		resData, err = data.Asset(binaryPath)
	} else if !s.FileExists(path) {
		err = fmt.Errorf("File does not exist at path: %s", path)
		log.Error(err)
	} else {
		resData, err = ioutil.ReadFile(path)
	}

	if err != nil {
		log.Errorf("Error reading file %s: %+v", path, err)
		return []byte{}, err
	}

	return resData, nil
}

// Walk walks the entire file structure for `root` and calls `walkFn` for each item it finds
func (s service) Walk(root string, walkFn WalkFunc) error {
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
func (s service) Append(targetFilePath string, content string) error {
	//Append second line
	file, err := os.OpenFile(targetFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, OwnerPermission)
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

// Write creates file at `targetFilePath` and appends the `content` to it. Deletes the file if it already exists
func (s service) Write(targetFilePath string, content string) error {
	s.DeletePath(targetFilePath)
	return s.Append(targetFilePath, content)
}

func (s service) IsValidPath(path string) error {
	_, err := os.Stat(path)
	return err
}

func (s service) Base64Encode(v []byte) (string, error) {
	return base64.StdEncoding.EncodeToString(v), nil
}
