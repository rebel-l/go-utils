package errorutils_test

import (
	"fmt"
	"testing"

	"github.com/rebel-l/go-utils/errorutils"
)

func TestEqual(t *testing.T) {
	testCases := []struct {
		name     string
		a        error
		b        error
		expected bool
	}{
		{
			name:     "a & b nil",
			expected: true,
		},
		{
			name:     "a nil, b not nil",
			b:        fmt.Errorf("something"),
			expected: false,
		},
		{
			name:     "a not, b nil",
			a:        fmt.Errorf("something else"),
			expected: false,
		},
		{
			name:     "a not nil, b not nil and same message",
			a:        fmt.Errorf("something"),
			b:        fmt.Errorf("something"),
			expected: true,
		},
		{
			name:     "a not nil, b not nil and not same message",
			a:        fmt.Errorf("something"),
			b:        fmt.Errorf("something else"),
			expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := errorutils.Equal(testCase.a, testCase.b)
			if got != testCase.expected {
				t.Errorf("Expected %t but got %t", testCase.expected, got)
			}
		})
	}
}
