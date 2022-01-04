package httputils_test

import (
	"testing"
	"time"

	"github.com/rebel-l/go-utils/httputils"
)

func TestNewClient(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		option   interface{}
		expected time.Duration
	}{
		{
			name:     "default - no option",
			expected: httputils.ClientDefaultTimeout,
		},
		{
			name:     "option - time.Duration",
			option:   10 * time.Second,
			expected: 10 * time.Second,
		},
		{
			name:     "option - int",
			option:   30,
			expected: 30 * time.Second,
		},
	}

	for _, testCase := range testCases {
		tc := testCase

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			client := httputils.NewClient(tc.option)
			if tc.expected != client.Timeout {
				t.Errorf("expected timeout %v but got %v", tc.expected, client.Timeout)
			}
		})
	}
}
