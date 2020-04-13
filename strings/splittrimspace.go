package strings

import "strings"

// SplitTrimSpace combines strings.Split() and strings.TrimSpace() to return a slice of strings without spaces in the
// elements.
func SplitTrimSpace(s, sep string) []string {
	result := strings.Split(s, sep)

	for k, v := range result {
		result[k] = strings.TrimSpace(v)
	}

	return result
}
