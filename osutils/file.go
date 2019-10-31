package osutils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// FileOrPathExists checks if a path or file exists
func FileOrPathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}

// CopyFile copies a file from source to destination
func CopyFile(src, dest string) error {
	src = filepath.Clean(src)
	dest = filepath.Clean(dest)

	from, err := os.Open(src) // nolint: gosec
	if err != nil {
		return fmt.Errorf("copy failed on source file: %s", err)
	}
	defer func() { // nolint: wsl
		_ = from.Close() // nolint: gosec
	}()

	to, err := os.OpenFile(dest, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("copy failed on destination file: %s", err)
	}
	defer func() { // nolint: wsl
		_ = to.Close() // nolint: gosec
	}()

	if _, err = io.Copy(to, from); err != nil {
		return fmt.Errorf("copy failed: %s", err)
	}

	return nil
}
