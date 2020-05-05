package osutils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const (
	defaultPermissionDirectory os.FileMode = 0755
	defaultPermissionFile      os.FileMode = 0644
)

var (
	// ErrCopyFailed defines the error if copy operation failed.
	ErrCopyFailed = fmt.Errorf("copy failed")
)

// FileOrPathExists checks if a path or file exists.
func FileOrPathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}

// CopyFile copies a file from source to destination.
func CopyFile(src, dest string) error {
	src = filepath.Clean(src)
	dest = filepath.Clean(dest)

	from, err := os.Open(src) // nolint: gosec
	if err != nil {
		return fmt.Errorf("%w on source file: %s", ErrCopyFailed, err)
	}
	defer func() { // nolint: wsl
		_ = from.Close()
	}()

	to, err := os.OpenFile(dest, os.O_RDWR|os.O_CREATE, defaultPermissionFile)
	if err != nil {
		return fmt.Errorf("%w on destination file: %s", ErrCopyFailed, err)
	}
	defer func() { // nolint: wsl
		_ = to.Close()
	}()

	if _, err = io.Copy(to, from); err != nil {
		return fmt.Errorf("%w: %s", ErrCopyFailed, err)
	}

	return nil
}

// CreateDirectoryIfNotExists creates a path recursive.
func CreateDirectoryIfNotExists(path string, permission ...os.FileMode) (err error) {
	perm := defaultPermissionDirectory

	if len(permission) > 0 {
		perm = permission[0]
	}

	exist := FileOrPathExists(path)
	if !exist {
		err = os.MkdirAll(path, perm)
	}

	return
}

// CreateFileIfNotExists creates a file with the given file name (needs to include the path).
// If path doesn'T exist it returns an error.
func CreateFileIfNotExists(file string) (err error) {
	exist := FileOrPathExists(file)
	if !exist {
		_, err = os.Create(file)
	}

	return
}
