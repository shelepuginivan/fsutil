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

func TestFirstExistingDir(t *testing.T) {
	t.Run("should return first existing directory", func(t *testing.T) {
		tmp := t.TempDir()

		f1 := filepath.Join(tmp, "file1")
		os.WriteFile(f1, []byte("1"), os.ModePerm)

		f2 := filepath.Join(tmp, "file2")
		os.WriteFile(f2, []byte("2"), os.ModePerm)

		f3 := filepath.Join(tmp, "file3")
		os.WriteFile(f3, []byte("3"), os.ModePerm)

		d1 := t.TempDir()

		first, exists := fsutil.FirstExistingDir(f1, f2, tmp, f3, d1)
		assert.Equal(t, tmp, first)
		assert.True(t, exists)
	})

	t.Run("should return false if there are no existing directories", func(t *testing.T) {
		tmp := t.TempDir()

		f1 := filepath.Join(tmp, "file1")
		os.WriteFile(f1, []byte("1"), os.ModePerm)

		f2 := filepath.Join(tmp, "file2")
		os.WriteFile(f2, []byte("2"), os.ModePerm)

		f3 := filepath.Join(tmp, "file3")
		os.WriteFile(f3, []byte("3"), os.ModePerm)

		first, exists := fsutil.FirstExistingDir(f1, f2, f3)
		assert.Empty(t, first)
		assert.False(t, exists)
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

func TestFirstExistingFile(t *testing.T) {
	t.Run("should return first existing file", func(t *testing.T) {
		tmp := t.TempDir()

		p1 := filepath.Join(tmp, "file1")
		os.WriteFile(p1, []byte("1"), os.ModePerm)

		p2 := filepath.Join(tmp, "file2")
		os.WriteFile(p2, []byte("2"), os.ModePerm)

		d1 := t.TempDir()
		d2 := t.TempDir()

		first, exists := fsutil.FirstExistingFile(tmp, d1, p1, p2, d2)
		assert.Equal(t, p1, first)
		assert.True(t, exists)
	})

	t.Run("should return false if there are no existing files", func(t *testing.T) {
		d1 := t.TempDir()
		d2 := t.TempDir()
		d3 := t.TempDir()

		first, exists := fsutil.FirstExistingFile(d1, d2, d3)
		assert.Empty(t, first)
		assert.False(t, exists)
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

func TestFirstExistingPath(t *testing.T) {
	t.Run("should return first existing path", func(t *testing.T) {
		tmp := t.TempDir()

		p1 := filepath.Join(tmp, "file1")
		os.WriteFile(p1, []byte("1"), os.ModePerm)

		p2 := filepath.Join(tmp, "file2")
		os.WriteFile(p2, []byte("2"), os.ModePerm)

		d1 := t.TempDir()
		d2 := t.TempDir()

		first, exists := fsutil.FirstExistingPath(tmp, d1, p1, p2, d2)
		assert.Equal(t, tmp, first)
		assert.True(t, exists)
	})

	t.Run("should return false if there are no existing paths", func(t *testing.T) {
		first, exists := fsutil.FirstExistingPath(
			"/tmp/this/path/does/not/exist",
			"/owr24yutr084yut98r23yt4",
			"f20ifh438fh34rf34",
		)

		assert.Empty(t, first)
		assert.False(t, exists)
	})
}
