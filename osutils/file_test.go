package osutils_test

import (
	"io/ioutil"
	"os"
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

func TestCopyFile_Happy(t *testing.T) {
	source := "./../LICENSE"
	destination := "./../tmp/LICENSE"

	err := osutils.CopyFile(source, destination)
	if err != nil {
		t.Fatalf("expected no error on copying but got %s", err)
	}

	content, err := ioutil.ReadFile(destination)
	if err != nil {
		t.Fatalf("expected no error on laoding destination file but got %s", err)
	}

	if len(content) < 100 {
		t.Errorf("expected content to be more than %d cahraters but got only %d", 100, len(content))
	}

	if err = os.Remove(destination); err != nil {
		t.Fatalf("unable to cleanup after test execution: %s", err)
	}
}

func TestCopyFile_Unhappy(t *testing.T) {
	testCases := []struct {
		name        string
		src         string
		dest        string
		expectedErr string
	}{
		{
			name:        "wrong source",
			src:         "file_does_not_exist.log",
			dest:        "/tmp/something.log",
			expectedErr: "copy failed on source file: open file_does_not_exist.log: The system cannot find the file specified.",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := osutils.CopyFile(testCase.src, testCase.dest)
			if err == nil {
				t.Fatalf("excpected error is happen but got none")
			}

			if testCase.expectedErr != err.Error() {
				t.Errorf("expected error message '%s' but got '%s'", testCase.expectedErr, err.Error())
			}
		})
	}
}
