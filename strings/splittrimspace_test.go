package strings_test

import (
	"testing"

	"github.com/rebel-l/go-utils/strings"
)

func TestSplitTrimSpace(t *testing.T) {
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
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := strings.SplitTrimSpace(testCase.testString, ",")

			if len(testCase.expected) != len(actual) {
				t.Fatalf("expected %d elements but got %d", len(testCase.expected), len(actual))
			}

			for k, expected := range testCase.expected {
				if expected != actual[k] {
					t.Fatalf("expected elemet %d to be '%s' but got '%s'", k, expected, actual[k])
				}
			}
		})
	}
}
