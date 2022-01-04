package testingutils

import "time"

// TimeParse parses a given time string to the defined layout and suppresses the error from time.Parse().
// NOTE: Don't use this for production code, it's just a helper to generate quickly time.Time structs in unit tests.
func TimeParse(layout, value string) time.Time {
	t, _ := time.Parse(layout, value)

	return t
}
