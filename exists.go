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

// FirstExistingDir returns the first path that exists and is a directory.
// The second value returned reports whether at least one path exists.
func FirstExistingDir(paths ...string) (string, bool) {
	for _, path := range paths {
		if DirExists(path) {
			return path, true
		}
	}
	return "", false
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

// FirstExistingFile returns the first path that exists and is a regular file.
// The second value returned reports whether at least one path exists.
func FirstExistingFile(paths ...string) (string, bool) {
	for _, path := range paths {
		if FileExists(path) {
			return path, true
		}
	}
	return "", false
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

// FirstExistingPath returns the first path that exists.
// The second value returned reports whether at least one path exists.
func FirstExistingPath(paths ...string) (string, bool) {
	for _, path := range paths {
		if PathExists(path) {
			return path, true
		}
	}
	return "", false
}
