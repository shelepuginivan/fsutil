package fsutil_test

import (
	"testing"

	"github.com/shelepuginivan/fsutil"
	"github.com/stretchr/testify/assert"
)

func TestReplaceForbiddenChars(t *testing.T) {
	cases := []struct {
		name     string
		expected string
	}{
		{name: "><:\"\\/|?*", expected: "---------"},
		{name: ">>>name<<<", expected: "---name---"},
		{name: ">", expected: "-"},
		{name: "<", expected: "-"},
		{name: ":", expected: "-"},
		{name: "\"", expected: "-"},
		{name: "\\", expected: "-"},
		{name: "/", expected: "-"},
		{name: "|", expected: "-"},
		{name: "?", expected: "-"},
		{name: "*", expected: "-"},
		{name: " name ", expected: "name"},
		{name: "a_valid_filename.txt", expected: "a_valid_filename.txt"},
		{name: "", expected: ""},
	}

	for _, c := range cases {
		actual := fsutil.ReplaceForbiddenChars(c.name)

		assert.Equal(t, c.expected, actual)
	}
}

func TestReplaceForbiddenCharsWith(t *testing.T) {
	cases := []struct {
		name     string
		newchar  string
		expected string
	}{
		{name: "><:\"\\/|?*", newchar: "-", expected: "---------"},
		{name: ">>>name<<<", newchar: "-", expected: "---name---"},
		{name: ">", newchar: "", expected: ""},
		{name: "<", newchar: "-", expected: "-"},
		{name: ":", newchar: "-", expected: "-"},
		{name: "\"", newchar: "g", expected: "g"},
		{name: "\\", newchar: "-", expected: "-"},
		{name: "/", newchar: "-", expected: "-"},
		{name: "|", newchar: "-", expected: "-"},
		{name: "?_?", newchar: "multichar", expected: "multichar_multichar"},
		{name: "*", newchar: "-", expected: "-"},
		{name: " name ", newchar: "-", expected: "name"},
		{name: "a_valid_filename.txt", newchar: "-", expected: "a_valid_filename.txt"},
		{name: "", newchar: "-", expected: ""},
	}

	for _, c := range cases {
		actual := fsutil.ReplaceForbiddenCharsWith(c.name, c.newchar)

		assert.Equal(t, c.expected, actual)
	}
}
