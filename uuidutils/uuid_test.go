package uuidutils_test

import (
	"testing"

	"github.com/rebel-l/go-utils/uuidutils"

	"github.com/google/uuid"
)

func TestIsEmpty(t *testing.T) {
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
		t.Run(testCase.name, func(t *testing.T) {
			got := uuidutils.IsEmpty(testCase.uuid)
			if testCase.expected != got {
				t.Errorf("expected %t but got %t", testCase.expected, got)
			}
		})
	}
}
