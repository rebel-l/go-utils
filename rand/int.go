package rand

import (
	"math/rand"
)

const (
	defaultSummand = 1
)

// Int returns a random integer number between a min and max value.
func Int(min, max int) int {
	InitSeed()

	return rand.Intn(max-min+defaultSummand) + min
}
