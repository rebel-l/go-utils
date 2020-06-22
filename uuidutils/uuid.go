package uuidutils

import (
	"fmt"
)

const (
	// EmptyUUID is the string representation of an empty UUID.
	EmptyUUID = "00000000-0000-0000-0000-000000000000"
)

// IsEmpty returns true if the given UUID is empty.
func IsEmpty(u fmt.Stringer) bool {
	return u.String() == EmptyUUID
}
