package httputils

import (
	"net/http"
	"time"
)

const (
	// ClientDefaultTimeout is the default timeout for the http.Client.
	ClientDefaultTimeout = 5 * time.Second
)

// NewClient return a new http.Client. By default the timeout of the client is defined by the
// constant ClientDefaultTimeout.
// With the options parameter you you can change the timeout. Options accept a time.Duration or int which means seconds.
func NewClient(options ...interface{}) *http.Client {
	timeout := ClientDefaultTimeout

	if len(options) > 0 {
		for _, v := range options {
			switch x := v.(type) {
			case time.Duration:
				timeout = x
			case int:
				timeout = time.Duration(x) * time.Second
			}
		}
	}

	return &http.Client{
		Timeout: timeout,
	}
}
