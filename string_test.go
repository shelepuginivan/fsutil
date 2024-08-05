package fsutil_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/shelepuginivan/fsutil"
	"github.com/stretchr/testify/assert"
)

func TestReadString(t *testing.T) {
	t.Run("should read file content as string", func(t *testing.T) {
		expected := "lorem ipsum dolor sit amet ..."
		path := filepath.Join(t.TempDir(), "file.txt")

		os.WriteFile(path, []byte(expected), os.ModePerm)

		actual, err := fsutil.ReadString(path)

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("should return an empty string and an error if file cannot be read", func(t *testing.T) {
		actual, err := fsutil.ReadString("this does not exist.")

		assert.Error(t, err)
		assert.Equal(t, "", actual)
	})
}

func TestReadLines(t *testing.T) {
	t.Run("should return content as a slice of string", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "file.txt")
		expected := []string{
			"Maybe it's a dream, maybe nothing else is real",
			"But it wouldn't mean a thing if I told you how I feel",
			"So I'm tired of all the pain, of the misery inside",
			"And I wish that I could live feeling nothing but the night",
			"You can tell me what to say, you can tell me where to go",
			"But I doubt that I would care, and my heart would never know",
			"If I make another move, there'll be no more turning back",
			"Because everything will change and it all will fade to black",
		}
		content := []byte(strings.Join(expected, "\n"))

		os.WriteFile(path, content, os.ModePerm)

		actual, err := fsutil.ReadLines(path)

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("should return error if file cannot be read", func(t *testing.T) {
		actual, err := fsutil.ReadLines("this does not exist.")

		assert.Error(t, err)
		assert.Nil(t, actual)
	})
}
