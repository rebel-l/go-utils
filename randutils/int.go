package randutils

import (
	"crypto/rand"
	"math/big"
)

const (
	defaultSummand = 1
)

// Int returns a random integer number between a min and max value.
func Int(min, max int) int {
	InitSeed()

	i, _ := rand.Int(rand.Reader, big.NewInt(int64(max-min+defaultSummand)))

	return int(i.Int64() + int64(min))
}
