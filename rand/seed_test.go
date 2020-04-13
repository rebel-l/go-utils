package rand

import "testing"

func TestSeedInitialized(t *testing.T) {
	actual := SeedInitialized()
	if actual {
		t.Errorf("expected that seed is not initialized at the beginning but got: %t", actual)
	}

	InitSeed()

	actual = SeedInitialized()
	if !actual {
		t.Errorf("expected that seed is initialized after call of InitSeed() but got: %t", actual)
	}
}
