/*
Package fsutil provides a collection of utility functions that extend the
capabilities of the Go standard filesystem (fs) package. It offers a variety of
tools for working with files and directories, making it easier to perform
common filesystem operations.

This package is designed to enhance the functionality of the standard library
by providing additional features and abstractions. It is particularly useful
for developers looking to simplify their filesystem interactions in Go
applications.
*/
package fsutil

// Must is a helper that wraps a call to a function returning an error and
// panics if the error is not nil. It is intended for use in variable
// initialization.
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
