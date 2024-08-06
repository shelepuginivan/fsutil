package fsutil_test

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/shelepuginivan/fsutil"
	"github.com/stretchr/testify/assert"
)

func TestCreateAll(t *testing.T) {
	t.Run("should create new file and parent directories", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "1", "2", "3", "file")
		file, err := fsutil.CreateAll(path)
		assert.NoError(t, err)
		file.Close()
	})

	t.Run("should return error if file cannot be created", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "1")
		filePath := filepath.Join(path, "2", "file")
		os.WriteFile(path, []byte("exists"), os.ModePerm)
		_, err := fsutil.CreateAll(filePath)
		assert.Error(t, err)
	})
}

func TestRemoveAllIfExists(t *testing.T) {
	t.Run("should remove all subdirectories and files of provided path", func(t *testing.T) {
		dir := t.TempDir()
		f1, _ := os.Create(filepath.Join(dir, "f1"))
		f2, _ := os.Create(filepath.Join(dir, "f1"))
		f3, _ := os.Create(filepath.Join(dir, "f1"))
		defer f1.Close()
		defer f2.Close()
		defer f3.Close()

		err := fsutil.RemoveAllIfExists(dir)
		assert.NoError(t, err)
		assert.NoDirExists(t, dir)
	})

	t.Run("should return error if path does not exist", func(t *testing.T) {
		assert.Error(t, fsutil.RemoveAllIfExists("does not exist"))
	})
}

func TestWriteAll(t *testing.T) {
	t.Run("should write file and create all parent directories", func(t *testing.T) {
		expected := []byte("content")
		path := filepath.Join(t.TempDir(), "1", "2", "3", "file")

		err := fsutil.WriteAll(path, expected)
		assert.NoError(t, err)

		actual, err := os.ReadFile(path)
		assert.NoError(t, err)
		assert.True(t, bytes.Equal(expected, actual))
	})

	t.Run("should return error if file cannot be written", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "1")
		filePath := filepath.Join(path, "2", "file")
		os.WriteFile(path, []byte("exists"), os.ModePerm)
		err := fsutil.WriteAll(filePath, []byte{})
		assert.Error(t, err)
	})
}
