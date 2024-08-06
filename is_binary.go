package fsutil

import (
	"bufio"
	"bytes"
	"unicode/utf8"
)

// IsBinary reports whether data is binary (non human-readable text).
// It checks the first and every 2^n line of the provided byte sequence
// (i.e. 0, 1, 2, 4, 8, ...), hence IsBinary may fail in rare cases.
//
// If data is an empty slice or nil, false is returned.
//
// For the most accurate determination, [IsBinarySlow] is recommended to be
// used.
// For the fastest determination, [IsBinaryFast] should be used.
func IsBinary(data []byte) bool {
	if len(data) == 0 || data == nil {
		return false
	}

	buf := bytes.NewBuffer(data)

	// NOTE: The split function defaults to ScanLines.
	scanner := bufio.NewScanner(buf)

	// Which lines to check, i.e. 0, 1, 2, 4, 8, ...
	next := 0

	for i := 0; scanner.Scan(); i++ {
		if i != next {
			continue
		}

		if !utf8.Valid(scanner.Bytes()) {
			return true
		}

		if next == 0 {
			next++
		} else {
			next *= 2
		}
	}

	return false
}

// IsBinarySlow reports whether data is binary (non human-readable text).
// It checks the entire file, hence the determination accuracy is the highest
// at a cost of performance.
//
// If data is an empty slice or nil, false is returned.
//
// For the most cases, [IsBinary] is recommended to be used.
// For the fastest determination, [IsBinaryFast] is recommended be used.
func IsBinarySlow(data []byte) bool {
	if len(data) == 0 || data == nil {
		return false
	}
	return !utf8.Valid(data)
}

// IsBinarySlow reports whether data is binary (non human-readable text). It
// checks only first line of the file, hence the determination performance is
// the highest at a cost of accuracy.
//
// If data is an empty slice or nil, false is returned.
//
// For the most cases, [IsBinary] is recommended to be used.
// For the most accurate determination, [IsBinarySlow] is recommended to be
// used.
func IsBinaryFast(data []byte) bool {
	if len(data) == 0 || data == nil {
		return false
	}
	buf := bytes.NewBuffer(data)

	// NOTE: The split function defaults to ScanLines.
	scanner := bufio.NewScanner(buf)
	scanner.Scan()

	return !utf8.Valid(scanner.Bytes())
}
