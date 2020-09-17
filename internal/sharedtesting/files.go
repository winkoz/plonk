package sharedtesting

import (
	"os"

	"github.com/winkoz/plonk/internal/io/log"
)

// CreatePath in the local filesystem
func CreatePath(path string) error {
	err := os.MkdirAll(path, 0755)
	if !os.IsExist(err) {
		log.Error(err)
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
