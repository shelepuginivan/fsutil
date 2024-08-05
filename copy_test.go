package fsutil_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/shelepuginivan/fsutil"
	"github.com/stretchr/testify/assert"
)

func TestCopyFile(t *testing.T) {
	t.Run("should copy a file", func(t *testing.T) {
		dir := t.TempDir()
		src := filepath.Join(dir, "src.txt")
		dst := filepath.Join(dir, "dst.txt")

		expected := []byte("some content that is about to be copied")
		os.WriteFile(src, expected, os.ModePerm)

		err := fsutil.CopyFile(src, dst)
		assert.NoError(t, err)

		actual, err := os.ReadFile(dst)
		assert.NoError(t, err) // Assert that file exists.
		assert.Equal(t, expected, actual)
	})

	t.Run("should return error if file cannot be copied", func(t *testing.T) {
		err := fsutil.CopyFile("this does not exist", "so this is not created")
		assert.Error(t, err)
	})
}
