// Package errorutils provides functions to handle Gos error interface
package errorutils

// Equal tests if two errors equals
func Equal(a, b error) bool {
	if a == nil && b == nil {
		return true
	}

	if a != nil && b != nil && a.Error() == b.Error() {
		return true
	}

	return false
}
