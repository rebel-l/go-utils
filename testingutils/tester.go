package testingutils

// Tester is the interface for go tests.
type Tester interface {
	Error(args ...interface{})
	Fatal(args ...interface{})
	Helper()
}
