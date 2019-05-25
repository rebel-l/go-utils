package option_test

import (
	"fmt"
	"testing"

	"github.com/rebel-l/go-utils/errorutils"

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
				{
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
				{
					Key:         "key1",
					Description: "description1",
				},
				{
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
				{
					Key:         "mykey",
					Description: "mydescription",
				},
			},
			key:      "mykey",
			expected: true,
		},
		{
			name: "one entry - find key case sensitive",
			data: option.Options{
				{
					Key:         "MyKey",
					Description: "mydescription",
				},
			},
			key:      "mykey",
			expected: false,
		},
		{
			name: "two entries - find key as second",
			data: option.Options{
				{
					Key:         "key1",
					Description: "description1",
				},
				{
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

func TestOptions_IsValidOptionCI(t *testing.T) {
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
				{
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
				{
					Key:         "key1",
					Description: "description1",
				},
				{
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
				{
					Key:         "mykey",
					Description: "mydescription",
				},
			},
			key:      "mykey",
			expected: true,
		},
		{
			name: "one entry - find key case insensitive",
			data: option.Options{
				{
					Key:         "MyKey",
					Description: "mydescription",
				},
			},
			key:      "mykey",
			expected: true,
		},
		{
			name: "two entries - find key as second",
			data: option.Options{
				{
					Key:         "key1",
					Description: "description1",
				},
				{
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
			got := testCase.data.IsValidOptionCI(testCase.key)
			if got != testCase.expected {
				t.Errorf("Expected %t, but got %t", testCase.expected, got)
			}
		})
	}
}

func TestOptions_ForAll(t *testing.T) {
	testCases := []struct {
		name     string
		data     option.Options
		callback func(option option.Option) error
		expected error
	}{
		{
			name: "empty - no error",
			callback: func(option option.Option) error {
				return nil
			},
		},
		{
			name: "empty - error",
			callback: func(option option.Option) error {
				return fmt.Errorf("failed")
			},
		},
		{
			name: "one entry - no error",
			data: option.Options{
				{
					Key:         "something",
					Description: "something is nothing",
				},
			},
			callback: func(option option.Option) error {
				return nil
			},
		},
		{
			name: "two entries - no error",
			data: option.Options{
				{
					Key:         "something1",
					Description: "something will be nothing",
				},
				{
					Key:         "something2",
					Description: "something is now nothing",
				},
			},
			callback: func(option option.Option) error {
				return nil
			},
		},
		{
			name: "one entry - cause error",
			data: option.Options{
				{
					Key:         "something",
					Description: "something is nothing",
				},
			},
			callback: func(option option.Option) error {
				return fmt.Errorf("failed")
			},
			expected: fmt.Errorf("failed to execute callback on entry 0: failed"),
		},
		{
			name: "two entries - second causes error",
			data: option.Options{
				{
					Key:         "something1",
					Description: "something will be nothing",
				},
				{
					Key:         "something2",
					Description: "something is now nothing",
				},
			},
			callback: func(option option.Option) error {
				if option.Key == "something2" {
					return fmt.Errorf("failed")
				}
				return nil
			},
			expected: fmt.Errorf("failed to execute callback on entry 1: failed"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := testCase.data.ForAll(testCase.callback)

			if !errorutils.Equal(got, testCase.expected) {
				t.Errorf("Expected result from callback '%s' but got '%s'", testCase.expected, got)
			}
		})
	}
}
