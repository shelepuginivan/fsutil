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

// WriteLines writes a slice of string as lines to a file.
//
// WriteLines uses platform-specific linebreak character (LF on Unix, CRLF of Windows).
// Use [WriteLinesWithNewlineChar] if specific newline character is required.
//
// Permission perm is optional and defaults to [os.ModePerm].
func WriteLines(path string, content []string, perm ...fs.FileMode) error {
	return WriteLinesWithNewlineChar(path, content, NewLine, perm...)
}

// WriteLinesWithNewlineChar is like [WriteLines], but allows to set newline character used.
func WriteLinesWithNewlineChar(path string, content []string, newline string, perm ...fs.FileMode) error {
	// Append the default permission.
	perm = append(perm, os.ModePerm)

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, perm[0])
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for i, line := range content {
		writer.WriteString(line)

		// Write a newline character if this is not the last line.
		if i+1 < len(content) {
			writer.WriteString(NewLine)
		}
	}

	return writer.Flush()
}
