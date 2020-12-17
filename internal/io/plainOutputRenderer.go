package io

import (
	"log"
	"os"
)

// PlainOutputRenderer renders bytes output as simple strings to console
type PlainOutputRenderer struct {
	log *log.Logger
}

// NewPlainOutputRenderer creates a plain output renderer that renders buffers into the console via stdout
func NewPlainOutputRenderer() PlainOutputRenderer {
	return PlainOutputRenderer{
		log: log.New(os.Stdout, "", 0),
	}
}

// RenderComponents renders the passed in output to the console via simple stdout call
func (pr PlainOutputRenderer) RenderComponents(output []byte) {
	pr.log.Print(string(output))
}

// RenderLogs renders the passed in output to the console via simple stdout call
func (pr PlainOutputRenderer) RenderLogs(output []byte) {
	pr.log.Print(string(output))
}
