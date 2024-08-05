package fsutil_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/shelepuginivan/fsutil"
	"github.com/stretchr/testify/assert"
)

func TestMust(t *testing.T) {
	t.Run("should not panic and return value if error is nil", func(t *testing.T) {
		fsutil.Must(os.ReadDir(t.TempDir()))
	})

	t.Run("should panic if error is not nil", func(t *testing.T) {
		assert.Panics(t, func() {
			fsutil.Must(func() (string, error) {
				return "", fmt.Errorf("this error is not nil, should panic")
			}())
		})
	})
}
