package osutils_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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

	if len(content) < 100 { //nolint: gomnd
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
			expectedErr: "copy failed on source file: open file_does_not_exist.log",
		},
		{
			name:        "wrong destination",
			src:         "./../LICENSE",
			dest:        "./../wrong_path/file_does_not_exist.log",
			expectedErr: "copy failed on destination file: open",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := osutils.CopyFile(testCase.src, testCase.dest)
			if err == nil {
				t.Fatalf("excpected error is happen but got none")
			}

			if !strings.Contains(err.Error(), testCase.expectedErr) {
				t.Errorf("expected error message contains '%s'. Got '%s'", testCase.expectedErr, err.Error())
			}
		})
	}
}

func TestCreateDirectoryIfNotExists(t *testing.T) {
	testCases := []struct {
		name           string
		path           string
		iterations     int
		levels         int
		withPermissons os.FileMode
		expectedError  error
	}{
		{
			name:       "happy - new directory one level",
			path:       "./../tmp/new",
			levels:     1, //nolint: gomnd
			iterations: 1, //nolint: gomnd
		},
		{
			name:       "happy - new directory two levels",
			path:       "./../tmp/new/two",
			levels:     2, //nolint: gomnd
			iterations: 1, //nolint: gomnd
		},
		{
			name:       "happy - new directory which already exists",
			path:       "./../tmp/new",
			levels:     1, //nolint: gomnd
			iterations: 2, //nolint: gomnd
		},
		{
			name:           "happy - new directory one level with permissions",
			path:           "./../tmp/new-withpermissions",
			withPermissons: 0777, //nolint: gomnd
			levels:         1,    //nolint: gomnd
			iterations:     1,    //nolint: gomnd
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			for i := 0; i < testCase.iterations; i++ {
				var err error

				if testCase.withPermissons > 0 {
					err = osutils.CreateDirectoryIfNotExists(testCase.path, testCase.withPermissons)
				} else {
					err = osutils.CreateDirectoryIfNotExists(testCase.path)
				}

				if err != testCase.expectedError {
					t.Fatalf("failed to create directory %s: expected error %v but got %v", testCase.path, testCase.expectedError, err)
				}
			}

			pathToCleanUp := testCase.path

			for i := 0; i < testCase.levels; i++ {
				if err := os.Remove(pathToCleanUp); err != nil {
					t.Fatalf("unable to cleanup after test execution: %s", err)
				}

				pathToCleanUp = filepath.Dir(pathToCleanUp)
			}
		})
	}
}

func TestCreateFileIfNotExists(t *testing.T) {
	testCases := []struct {
		name          string
		fileName      string
		iterations    int
		expectedError string
	}{
		{
			name:       "happy - new file",
			fileName:   "./../tmp/new_file.log",
			iterations: 1, //nolint: gomnd
		},
		{
			name:       "happy - file exists",
			fileName:   "./../tmp/file_exists.log",
			iterations: 2, //nolint: gomnd
		},
		{
			name:          "unhappy - path doesn't exist",
			fileName:      "./../tmp/pathdoesnotexist/new_file.log",
			iterations:    1,                                              //nolint: gomnd
			expectedError: "open ./../tmp/pathdoesnotexist/new_file.log:", // nolint: golint
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			for i := 0; i < testCase.iterations; i++ {
				err := osutils.CreateFileIfNotExists(testCase.fileName)
				if err != nil {
					if testCase.expectedError != "" && !strings.Contains(err.Error(), testCase.expectedError) {
						t.Fatalf(
							"failed to create file %s: expected error '%v' but got '%v'",
							testCase.fileName,
							testCase.expectedError,
							err,
						)
					}
					return
				}
			}

			if err := os.Remove(testCase.fileName); err != nil {
				t.Fatalf("unable to cleanup after test execution: %s", err)
			}
		})
	}
}
