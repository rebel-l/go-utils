package array_test

import (
	"testing"

	"github.com/rebel-l/go-utils/array"
)

func TestStringArrayEquals(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		a        []string
		b        []string
		expected bool
	}{
		{
			name:     "empty arrays",
			expected: true,
		},
		{
			name:     "one element same",
			a:        []string{"a"},
			b:        []string{"a"},
			expected: true,
		},
		{
			name:     "two elements same",
			a:        []string{"a", "b"},
			b:        []string{"a", "b"},
			expected: true,
		},
		{
			name:     "different length",
			a:        []string{"a"},
			b:        []string{"b", "c"},
			expected: false,
		},
		{
			name:     "one element different",
			a:        []string{"a"},
			b:        []string{"b"},
			expected: false,
		},
		{
			name:     "two elements different order",
			a:        []string{"a", "b"},
			b:        []string{"b", "a"},
			expected: false,
		},
		{
			name:     "two elements different values",
			a:        []string{"a", "b"},
			b:        []string{"c", "d"},
			expected: false,
		},
	}

	for _, testCase := range testCases {
		tc := testCase

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := array.StringArrayEquals(tc.a, tc.b)
			if tc.expected != actual {
				t.Errorf("Expected that array %v & %v equals '%t' but got '%t'", tc.a, tc.b, tc.expected, actual)
			}
		})
	}
}
