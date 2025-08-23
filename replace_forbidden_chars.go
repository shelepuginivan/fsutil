package fsutil

import "strings"

// ReplaceForbiddenChars replaces characters forbidden in some filesystems with
// dash (-).
//
// Characters replaced are:
//   - < (less than)
//   - > (greater than)
//   - : (colon)
//   - " (double quote)
//   - / (slash)
//   - \ (backslash)
//   - | (pipe)
//   - ? (question mark)
//   - * (asterisk)
//
// In addition, all leading and trailing whitespace characters are trimmed.
func ReplaceForbiddenChars(name string) string {
	return ReplaceForbiddenCharsWith(name, "-")
}

// ReplaceForbiddenCharsWith replaces characters forbidden in some filesystems
// with the specified character.
//
// Characters replaced are:
//   - < (less than)
//   - > (greater than)
//   - : (colon)
//   - " (double quote)
//   - / (slash)
//   - \ (backslash)
//   - | (pipe)
//   - ? (question mark)
//   - * (asterisk)
//
// In addition, all leading and trailing whitespace characters are trimmed.
func ReplaceForbiddenCharsWith(name string, newchar string) string {
	r := strings.NewReplacer(
		"<", newchar,
		">", newchar,
		":", newchar,
		"\"", newchar,
		"/", newchar,
		"\\", newchar,
		"|", newchar,
		"?", newchar,
		"*", newchar,
	)

	return strings.TrimSpace(r.Replace(name))
}
