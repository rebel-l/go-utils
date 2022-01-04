package uuidutils_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/rebel-l/go-utils/uuidutils"
)

func TestIsEmpty(t *testing.T) {
	t.Parallel()

	u, err := uuid.NewRandom()
	if err != nil {
		t.Fatalf("failed to generate UUID: %v", err)
	}

	testCases := []struct {
		name     string
		uuid     uuid.UUID
		expected bool
	}{
		{
			name:     "empty",
			expected: true,
		},
		{
			name:     "valid UUID",
			uuid:     u,
			expected: false,
		},
	}

	for _, testCase := range testCases {
		tc := testCase

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := uuidutils.IsEmpty(tc.uuid)
			if tc.expected != got {
				t.Errorf("expected %t but got %t", tc.expected, got)
			}
		})
	}
}
