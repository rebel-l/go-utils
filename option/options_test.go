package option_test

import (
	"testing"

	"github.com/rebel-l/go-utils/option"
)

func TestOptions_IsValidOption(t *testing.T) {
	testCases := []struct {
		name     string
		data     option.Options
		key      string
		expected bool
	}{
		{
			name:     "empty options",
			key:      "empty",
			expected: false,
		},
		{
			name: "one entry - miss key",
			data: option.Options{
				option.Option{
					Key:         "mykey",
					Description: "mydescription",
				},
			},
			key:      "miss",
			expected: false,
		},
		{
			name: "two entries - miss key",
			data: option.Options{
				option.Option{
					Key:         "key1",
					Description: "description1",
				},
				option.Option{
					Key:         "key2",
					Description: "description2",
				},
			},
			key:      "miss",
			expected: false,
		},
		{
			name: "one entry - find key",
			data: option.Options{
				option.Option{
					Key:         "mykey",
					Description: "mydescription",
				},
			},
			key:      "mykey",
			expected: true,
		},
		{
			name: "two entries - find key as second",
			data: option.Options{
				option.Option{
					Key:         "key1",
					Description: "description1",
				},
				option.Option{
					Key:         "key2",
					Description: "description2",
				},
			},
			key:      "key2",
			expected: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := testCase.data.IsValidOption(testCase.key)
			if got != testCase.expected {
				t.Errorf("Expected %t, but got %t", testCase.expected, got)
			}
		})
	}
}
