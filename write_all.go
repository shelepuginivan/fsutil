package fsutil

import (
	"os"
	"path/filepath"
)

// WriteAll writes data to path, creating all necessary parent directories.
// If parent directories already exist, [WriteAll] does nothing to them.
// If file already exists, it is overwritten with the given data.
func WriteAll(path string, data []byte) error {
	dir := filepath.Dir(path)

	err := os.MkdirAll(dir, os.ModePerm|os.ModeDir)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, os.ModePerm)
}
