package osutils_test

import (
	"testing"

	"github.com/rebel-l/go-utils/osutils"
)

func TestFileOrPathExists(t *testing.T) {
	testcases := []struct {
		name     string
		path     string
		expected bool
	}{
		{
			name:     "file exists",
			path:     "./file.go",
			expected: true,
		},
		{
			name:     "path exists",
			path:     "./../osutils",
			expected: true,
		},
		{
			name:     "file or path does not exists",
			path:     "./doesnotexist",
			expected: false,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			if actual := osutils.FileOrPathExists(testcase.path); testcase.expected != actual {
				t.Errorf("Expected result for existing files is %t but got %t", testcase.expected, actual)
			}
		})
	}
}
