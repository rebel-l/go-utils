package testingutils

import (
	"errors"
	"fmt"
)

// ErrorsCheck is a helper to assert two errors. Are save to test on nil.
func ErrorsCheck(t Tester, expected, actual error) {
	t.Helper()

	if errors.Is(actual, expected) {
		return
	}

	if expected != nil && actual != nil {
		if expected.Error() != actual.Error() {
			t.Error(fmt.Sprintf("expected error '%v' but got '%v'", expected, actual))
		}

		return
	}

	t.Error(fmt.Sprintf("expected error '%v' but got '%v'", expected, actual))
}
