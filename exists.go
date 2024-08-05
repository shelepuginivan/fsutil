package fsutil

import "os"

// DirExists reports whether path exists and is a directory.
func DirExists(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return stat.IsDir()
}

// FileExists reports whether path exists and is a file.
func FileExists(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !stat.IsDir()
}

// PathExists reports whether path exists.
func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
