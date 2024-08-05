package fsutil

import (
	"fmt"
	"os"
)

// DirExists reports whether path exists and is a directory.
func DirExists(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return stat.IsDir()
}

// AssertDirExists panics if path does not exist or is not a directory.
func AssertDirExists(path string) {
	if !DirExists(path) {
		panic(fmt.Errorf("%s does not exist or is not a directory", path))
	}
}

// FileExists reports whether path exists and is a file.
func FileExists(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !stat.IsDir()
}

// AssertFileExists panics if path does not exist or is not a file.
func AssertFileExists(path string) {
	if !FileExists(path) {
		panic(fmt.Errorf("%s does not exist or is not a file", path))
	}
}

// PathExists reports whether path exists.
func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// AssertPathExists panics if path does not exist.
func AssertPathExists(path string) {
	if !PathExists(path) {
		panic(fmt.Errorf("%s does not exist", path))
	}
}
