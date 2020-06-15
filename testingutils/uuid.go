package testingutils

import (
	"fmt"

	"github.com/google/uuid"
)

// Tester is the interface for go tests.
type Tester interface {
	Fatal(args ...interface{})
	Helper()
}

// UUIDParse is a test helper to generate UUID fixtures.
func UUIDParse(t Tester, s string) uuid.UUID {
	t.Helper()

	u, err := uuid.Parse(s)
	if err != nil {
		t.Fatal(fmt.Sprintf("failed to parse UUID from '%s': %v", s, err))
	}

	return u
}
