package fsutil_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/shelepuginivan/fsutil"
	"github.com/stretchr/testify/assert"
)

func TestDirExists(t *testing.T) {
	t.Run("should return true if directory exists", func(t *testing.T) {
		assert.True(t, fsutil.DirExists(t.TempDir()))
	})

	t.Run("should return false if directory does not exist", func(t *testing.T) {
		assert.False(t, fsutil.DirExists("this dir does not exist"))
	})

	t.Run("should return false if file exists, but is not a directory", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "file")
		os.WriteFile(path, []byte("file"), os.ModePerm)
		assert.False(t, fsutil.DirExists(path))
	})
}

func TestAssertDirExists(t *testing.T) {
	t.Run("should not panic if directory exists", func(t *testing.T) {
		assert.NotPanics(t, func() {
			fsutil.AssertDirExists(t.TempDir())
		})
	})

	t.Run("should panic if directory does not exist", func(t *testing.T) {
		assert.Panics(t, func() {
			fsutil.AssertDirExists("does not exist")
		})
	})

	t.Run("should panic if path exists but is not a directory", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "file")
		os.WriteFile(path, []byte("file"), os.ModePerm)
		assert.Panics(t, func() {
			fsutil.AssertDirExists(path)
		})
	})
}

func TestFileExists(t *testing.T) {
	t.Run("should return true if file exists", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "file")
		os.WriteFile(path, []byte("file"), os.ModePerm)
		assert.True(t, fsutil.FileExists(path))
	})

	t.Run("should return false if file does not exist", func(t *testing.T) {
		assert.False(t, fsutil.FileExists("this file does not exist"))
	})

	t.Run("should return false if file exists, but is not a file", func(t *testing.T) {
		assert.False(t, fsutil.FileExists(t.TempDir()))
	})
}

func TestAssertFileExists(t *testing.T) {
	t.Run("should not panic if file exists", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "file")
		os.WriteFile(path, []byte("file"), os.ModePerm)
		assert.NotPanics(t, func() {
			fsutil.AssertFileExists(path)
		})
	})

	t.Run("should panic if file does not exist", func(t *testing.T) {
		assert.Panics(t, func() {
			fsutil.AssertFileExists("does not exist")
		})
	})

	t.Run("should panic if path exists but is not a file", func(t *testing.T) {
		assert.Panics(t, func() {
			fsutil.AssertFileExists(t.TempDir())
		})
	})
}

func TestPathExists(t *testing.T) {
	t.Run("should return true if file exists", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "file")
		os.WriteFile(path, []byte("file"), os.ModePerm)
		assert.True(t, fsutil.PathExists(path))
	})

	t.Run("should return true if directory exists", func(t *testing.T) {
		assert.True(t, fsutil.PathExists(t.TempDir()))
	})

	t.Run("should return false if path does not exist", func(t *testing.T) {
		assert.False(t, fsutil.PathExists("this file does not exist"))
	})
}

func TestAssertPathExists(t *testing.T) {
	t.Run("should not panic if directory exists", func(t *testing.T) {
		assert.NotPanics(t, func() {
			fsutil.AssertPathExists(t.TempDir())
		})
	})

	t.Run("should not panic if file exists", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "file")
		os.WriteFile(path, []byte("file"), os.ModePerm)
		assert.NotPanics(t, func() {
			fsutil.AssertPathExists(path)
		})
	})

	t.Run("should panic if path does not exist", func(t *testing.T) {
		assert.Panics(t, func() {
			fsutil.AssertPathExists("does not exist")
		})
	})
}
