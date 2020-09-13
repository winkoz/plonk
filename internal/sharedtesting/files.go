package sharedtesting

import "os"

// CreatePath in the local filesystem
func CreatePath(path string) error {
	err := os.MkdirAll(path, 0755)
	if !os.IsExist(err) {
		return err
	}
	return nil
}

// DeletePath in the local filesystem
func DeletePath(path string) error {
	return os.RemoveAll(path)
}
