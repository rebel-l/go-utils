package randutils

import (
	"math/rand"
	"time"
)

var seeded = false // nolint: gochecknoglobals

// InitSeed ensures randutils package to return random values. Function ensures seed is initialized only once.
func InitSeed() {
	if SeedInitialized() {
		return
	}

	rand.Seed(time.Now().UnixNano())

	seeded = true
}

// SeedInitialized returns true if the seed was already initialized.
func SeedInitialized() bool {
	return seeded
}
