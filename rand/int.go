package rand

import (
	"math/rand"
	"time"
)

const (
	defaultSummand = 1
)

// Int returns a random integer number between a min and max value.
func Int(min, max int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max-min+defaultSummand) + min
}
