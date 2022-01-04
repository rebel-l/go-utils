package stringsutils_test

import (
	"testing"

	"github.com/rebel-l/go-utils/stringsutils"
)

func TestSplitTrimSpace(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name       string
		testString string
		expected   []string
	}{
		{
			name:       "one element without space",
			testString: "mystring",
			expected:   []string{"mystring"},
		},
		{
			name:       "one element with space",
			testString: " mystring ",
			expected:   []string{"mystring"},
		},
		{
			name:       "two elements without space",
			testString: "first,second",
			expected:   []string{"first", "second"},
		},
		{
			name:       "two elements with space",
			testString: " first, second ",
			expected:   []string{"first", "second"},
		},
		{
			name:       "empty string",
			testString: "",
			expected:   []string{},
		},
		{
			name:       "one empty element",
			testString: " ",
			expected:   []string{},
		},
		{
			name:       "two empty elements",
			testString: " , ",
			expected:   []string{},
		},
	}

	for _, testCase := range testCases {
		tc := testCase

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := stringsutils.SplitTrimSpace(tc.testString, ",")

			if len(tc.expected) != len(actual) {
				t.Fatalf("expected %d elements but got %d", len(tc.expected), len(actual))
			}

			for k, expected := range tc.expected {
				if expected != actual[k] {
					t.Fatalf("expected elemet %d to be '%s' but got '%s'", k, expected, actual[k])
				}
			}
		})
	}
}
