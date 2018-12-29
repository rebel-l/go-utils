package osutils

import "os"

// FileOrPathExists checks if a path or file exists
func FileOrPathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}
