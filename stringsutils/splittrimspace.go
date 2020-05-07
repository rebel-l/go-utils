package stringsutils

import "strings"

// SplitTrimSpace combines strings.Split() and strings.TrimSpace() to return a slice of strings without spaces in the
// elements. Empty elements are removed.
func SplitTrimSpace(s, sep string) []string {
	var result []string

	split := strings.Split(s, sep)

	for _, v := range split {
		e := strings.TrimSpace(v)
		if e != "" {
			result = append(result, e)
		}
	}

	return result
}
