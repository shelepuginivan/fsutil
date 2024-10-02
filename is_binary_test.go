package fsutil_test

import (
	"crypto/rand"
	"testing"
	"unicode/utf8"

	"github.com/shelepuginivan/fsutil"
	"github.com/stretchr/testify/assert"
)

// This was generated by 2goarray externally.
// https://github.com/cratonica/2goarray
var binarySlice []byte = []byte{
	0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a, 0x00, 0x00, 0x00, 0x0d,
	0x49, 0x48, 0x44, 0x52, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01,
	0x01, 0x03, 0x00, 0x00, 0x00, 0x25, 0xdb, 0x56, 0xca, 0x00, 0x00, 0x00,
	0x20, 0x63, 0x48, 0x52, 0x4d, 0x00, 0x00, 0x7a, 0x26, 0x00, 0x00, 0x80,
	0x84, 0x00, 0x00, 0xfa, 0x00, 0x00, 0x00, 0x80, 0xe8, 0x00, 0x00, 0x75,
	0x30, 0x00, 0x00, 0xea, 0x60, 0x00, 0x00, 0x3a, 0x98, 0x00, 0x00, 0x17,
	0x70, 0x9c, 0xba, 0x51, 0x3c, 0x00, 0x00, 0x00, 0x06, 0x50, 0x4c, 0x54,
	0x45, 0xec, 0xec, 0xeb, 0xff, 0xff, 0xff, 0x28, 0x05, 0xa8, 0xc3, 0x00,
	0x00, 0x00, 0x01, 0x62, 0x4b, 0x47, 0x44, 0x01, 0xff, 0x02, 0x2d, 0xde,
	0x00, 0x00, 0x00, 0x0a, 0x49, 0x44, 0x41, 0x54, 0x08, 0xd7, 0x63, 0x60,
	0x00, 0x00, 0x00, 0x02, 0x00, 0x01, 0xe2, 0x21, 0xbc, 0x33, 0x00, 0x00,
	0x00, 0x00, 0x49, 0x45, 0x4e, 0x44, 0xae, 0x42, 0x60, 0x82,
}

var textSlice []byte = []byte{
	0x49, 0x20, 0x77, 0x69, 0x6c, 0x6c, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x20, 0x69, 0x6e, 0x74, 0x6f, 0x20, 0x61, 0x20,
	0x6d, 0x61, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x20, 0x67, 0x69, 0x72, 0x6c,
	0x0a, 0x49, 0x20, 0x77, 0x69, 0x6c, 0x6c, 0x20, 0x64, 0x65, 0x66, 0x79,
	0x20, 0x6d, 0x79, 0x20, 0x6f, 0x77, 0x6e, 0x20, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x0a, 0x49, 0x20, 0x77, 0x69, 0x6c, 0x6c, 0x20, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x20, 0x69, 0x6e, 0x74, 0x6f,
	0x20, 0x61, 0x20, 0x6d, 0x61, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x20, 0x67,
	0x69, 0x72, 0x6c, 0x0a, 0x41, 0x6e, 0x64, 0x20, 0x6d, 0x79, 0x20, 0x76,
	0x6f, 0x64, 0x6b, 0x61, 0x20, 0x62, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x20,
	0x77, 0x69, 0x6c, 0x6c, 0x20, 0x74, 0x75, 0x72, 0x6e, 0x20, 0x69, 0x6e,
	0x74, 0x6f, 0x20, 0x61, 0x20, 0x6d, 0x61, 0x67, 0x69, 0x63, 0x20, 0x77,
	0x61, 0x6e, 0x64, 0x20, 0x27, 0x63, 0x61, 0x75, 0x73, 0x65, 0x0a, 0x49,
	0x20, 0x77, 0x69, 0x6c, 0x6c, 0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x6f, 0x72, 0x6d, 0x20, 0x69, 0x6e, 0x74, 0x6f, 0x20, 0x61, 0x20, 0x6d,
	0x61, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x20, 0x67, 0x69, 0x72, 0x6c, 0x0a,
	0x54, 0x68, 0x61, 0x74, 0x27, 0x73, 0x20, 0x68, 0x6f, 0x77, 0x20, 0x49,
	0xe2, 0x80, 0x99, 0x6c, 0x6c, 0x20, 0x6d, 0x65, 0x65, 0x74, 0x20, 0x79,
	0x6f, 0x75, 0x72, 0x20, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x0a, 0x49, 0x20, 0x77, 0x69, 0x6c, 0x6c, 0x20,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x20, 0x69, 0x6e,
	0x74, 0x6f, 0x20, 0x61, 0x20, 0x6d, 0x61, 0x67, 0x69, 0x63, 0x61, 0x6c,
	0x20, 0x67, 0x69, 0x72, 0x6c, 0x0a, 0x54, 0x68, 0x61, 0x74, 0x20, 0x77,
	0x69, 0x6c, 0x6c, 0x20, 0x67, 0x69, 0x76, 0x65, 0x20, 0x6d, 0x65, 0x20,
	0x61, 0x20, 0x63, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x20, 0x74, 0x6f, 0x20,
	0x72, 0x69, 0x73, 0x65, 0x0a,
}

func TestIsBinary(t *testing.T) {
	t.Run("should return true if the data is binary", func(t *testing.T) {
		assert.True(t, fsutil.IsBinary(binarySlice))
	})

	t.Run("should return false if the data is text", func(t *testing.T) {
		assert.False(t, fsutil.IsBinary(textSlice))
	})

	t.Run("should return false if the slice is empty or nil", func(t *testing.T) {
		assert.False(t, fsutil.IsBinary([]byte{}))
		assert.False(t, fsutil.IsBinary(nil))
	})

	t.Run("should return valid value (random bytes) at least in 99% cases", func(t *testing.T) {
		r := make([]byte, 10000)
		correct := 0
		minimum := 990

		for range 1000 {
			rand.Read(r)

			expected := !utf8.Valid(r)
			actual := fsutil.IsBinary(r)

			if expected == actual {
				correct += 1
			}
		}

		assert.GreaterOrEqual(t, correct, minimum)
	})
}

func TestIsBinarySlow(t *testing.T) {
	t.Run("should return true if the data is binary", func(t *testing.T) {
		assert.True(t, fsutil.IsBinarySlow(binarySlice))
	})

	t.Run("should return false if the data is text", func(t *testing.T) {
		assert.False(t, fsutil.IsBinarySlow(textSlice))
	})

	t.Run("should return false if the slice is empty or nil", func(t *testing.T) {
		assert.False(t, fsutil.IsBinarySlow([]byte{}))
		assert.False(t, fsutil.IsBinarySlow(nil))
	})

	t.Run("should return valid value (random bytes) in 100% cases", func(t *testing.T) {
		r := make([]byte, 10000)
		correct := 0
		minimum := 1000

		for range 1000 {
			rand.Read(r)

			expected := !utf8.Valid(r)
			actual := fsutil.IsBinarySlow(r)

			if expected == actual {
				correct += 1
			}
		}

		assert.GreaterOrEqual(t, correct, minimum)
	})
}

func TestIsBinaryFast(t *testing.T) {
	t.Run("should return true if the data is binary", func(t *testing.T) {
		assert.True(t, fsutil.IsBinaryFast(binarySlice))
	})

	t.Run("should return false if the data is text", func(t *testing.T) {
		assert.False(t, fsutil.IsBinaryFast(textSlice))
	})

	t.Run("should return false if the slice is empty or nil", func(t *testing.T) {
		assert.False(t, fsutil.IsBinaryFast([]byte{}))
		assert.False(t, fsutil.IsBinaryFast(nil))
	})

	t.Run("should return valid value (random bytes) at least in 95% cases", func(t *testing.T) {
		r := make([]byte, 10000)
		correct := 0
		minimum := 950

		for range 1000 {
			rand.Read(r)

			expected := !utf8.Valid(r)
			actual := fsutil.IsBinaryFast(r)

			if expected == actual {
				correct += 1
			}
		}

		assert.GreaterOrEqual(t, correct, minimum)
	})
}