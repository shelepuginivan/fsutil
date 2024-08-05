package fsutil

import (
	"bufio"
	"io/fs"
	"os"
)

// ReadString reads path and returns its content as string.
func ReadString(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// WriteString writes a string content to a file.
//
// Permission perm is optional and defaults to [os.ModePerm].
func WriteString(path, content string, perm ...fs.FileMode) error {
	// Append the default permission.
	perm = append(perm, os.ModePerm)
	return os.WriteFile(path, []byte(content), perm[0])
}

// ReadLines reads a file and returns its content as a slice of strings.
func ReadLines(path string) (lines []string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// The split function of [bufio.Scanner] defaults to SplitLines.
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
