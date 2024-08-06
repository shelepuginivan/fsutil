package fsutil

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// CreateAll creates file and all its parent directories.
// If the file exists, it is truncated.
func CreateAll(file string) (*os.File, error) {
	dir := filepath.Dir(file)

	if err := os.MkdirAll(dir, os.ModePerm|os.ModeDir); err != nil {
		return nil, err
	}

	return os.Create(file)
}

// RemoveAllIfExists is like [os.RemoveAll] except that it returns an error if
// the path does not exist.
func RemoveAllIfExists(path string) error {
	if !PathExists(path) {
		return fmt.Errorf("path %s does not exist", path)
	}
	return os.RemoveAll(path)
}

// WriteAll writes data to path, creating all necessary parent directories.
// If parent directories already exist, [WriteAll] does nothing to them.
// If file already exists, it is overwritten with the given data.
//
// Permission perm is optional and defaults to [os.ModePerm].
func WriteAll(path string, data []byte, perm ...fs.FileMode) error {
	// Append the default permission.
	perm = append(perm, os.ModePerm)

	dir := filepath.Dir(path)

	err := os.MkdirAll(dir, os.ModePerm|os.ModeDir)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, perm[0])
}
