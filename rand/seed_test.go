package rand_test

import (
	"testing"

	"github.com/rebel-l/go-utils/rand"
)

func TestSeedInitialized(t *testing.T) {
	actual := rand.SeedInitialized()
	if actual {
		t.Errorf("expected that seed is not initialized at the beginning but got: %t", actual)
	}

	rand.InitSeed()

	actual = rand.SeedInitialized()
	if !actual {
		t.Errorf("expected that seed is initialized after call of InitSeed() but got: %t", actual)
	}
}
