package fsutil

import "os"

// CopyFile copies file src to file dst. It does not support recursive copying.
func CopyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, data, os.ModePerm)
}
