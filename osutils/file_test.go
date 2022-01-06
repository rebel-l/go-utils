package osutils_test

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/rebel-l/go-utils/osutils"
)

func TestFileOrPathExists(t *testing.T) {
	t.Parallel()

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
		tc := testcase

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if actual := osutils.FileOrPathExists(tc.path); tc.expected != actual {
				t.Errorf("Expected result for existing files is %t but got %t", tc.expected, actual)
			}
		})
	}
}

func TestCopyFile_Happy(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

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
		tc := testCase

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := osutils.CopyFile(tc.src, tc.dest)
			if err == nil {
				t.Fatalf("excpected error is happen but got none")
			}

			if !strings.Contains(err.Error(), tc.expectedErr) {
				t.Errorf("expected error message contains '%s'. Got '%s'", tc.expectedErr, err.Error())
			}
		})
	}
}

func TestCreateDirectoryIfNotExists(t *testing.T) {
	t.Parallel()

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
			path:       "./../tmp/new1",
			levels:     1,
			iterations: 1,
		},
		{
			name:       "happy - new directory two levels",
			path:       "./../tmp/new2/two",
			levels:     2,
			iterations: 1,
		},
		{
			name:       "happy - new directory which already exists",
			path:       "./../tmp/new4",
			levels:     1,
			iterations: 2,
		},
		{
			name:           "happy - new directory one level with permissions",
			path:           "./../tmp/new-withpermissions",
			withPermissons: 0o777,
			levels:         1,
			iterations:     1,
		},
	}

	for _, testCase := range testCases {
		tc := testCase

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			for i := 0; i < tc.iterations; i++ {
				var err error

				if tc.withPermissons > 0 {
					err = osutils.CreateDirectoryIfNotExists(tc.path, tc.withPermissons)
				} else {
					err = osutils.CreateDirectoryIfNotExists(tc.path)
				}

				if !errors.Is(err, tc.expectedError) {
					t.Fatalf("failed to create directory %s: expected error %v but got %v", tc.path, tc.expectedError, err)
				}
			}

			pathToCleanUp := tc.path

			for i := 0; i < tc.levels; i++ {
				if err := os.Remove(pathToCleanUp); err != nil {
					t.Fatalf("unable to cleanup after test execution: %s", err)
				}

				pathToCleanUp = filepath.Dir(pathToCleanUp)
			}
		})
	}
}

func TestCreateFileIfNotExists(t *testing.T) {
	t.Parallel()

	// prepare
	path := "./../tmp/TestCreateFileIfNotExists"
	createdTestFiles := []string{
		path + "/new_file.log",
		path + "/file_exists.log",
	}

	if osutils.FileOrPathExists(path) {
		createdTestFiles = append(createdTestFiles, path)
		for _, v := range createdTestFiles {
			if err := os.Remove(v); err != nil && !strings.Contains(err.Error(), "The system cannot find the file specified.") {
				t.Fatalf("unable to cleanup test directory: %s", err)
			}
		}
	}

	if err := osutils.CreateDirectoryIfNotExists(path); err != nil {
		t.Fatalf("unable to create test directory: %s", err)
	}

	// test
	testCases := []struct {
		name          string
		fileName      string
		iterations    int
		expectedError string
	}{
		{
			name:       "happy - new file",
			fileName:   createdTestFiles[0],
			iterations: 1,
		},
		{
			name:       "happy - file exists",
			fileName:   createdTestFiles[1],
			iterations: 2,
		},
		{
			name:          "unhappy - path doesn't exist",
			fileName:      path + "/pathdoesnotexist/new_file.log",
			iterations:    1,
			expectedError: "open ./../tmp/TestCreateFileIfNotExists/pathdoesnotexist/new_file.log:",
		},
	}

	for _, testCase := range testCases {
		tc := testCase

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			for i := 0; i < tc.iterations; i++ {
				err := osutils.CreateFileIfNotExists(tc.fileName)
				if err != nil {
					if tc.expectedError != "" && !strings.Contains(err.Error(), tc.expectedError) {
						t.Fatalf(
							"failed to create file %s: expected error '%v' but got '%v'",
							tc.fileName,
							tc.expectedError,
							err,
						)
					}

					return
				}
			}
		})
	}
}
