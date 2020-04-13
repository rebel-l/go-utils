package rand_test

import (
	"testing"

	"github.com/rebel-l/go-utils/rand"
)

func TestInt(t *testing.T) {
	testCases := []struct {
		name string
		min  int
		max  int
	}{
		{
			name: "between 0 and 100",
			max:  100, //nolint: gomnd
		},
		{
			name: "between 20 and 50",
			min:  20, //nolint: gomnd
			max:  50, //nolint: gomnd
		},
		{
			name: "between -50 and 50",
			min:  -50, //nolint: gomnd
			max:  50,  //nolint: gomnd
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := rand.Int(testCase.min, testCase.max)

			if actual < testCase.min {
				t.Errorf("expected randon number greater or eqal %d, but got %d", testCase.min, actual)
			}

			if actual > testCase.max {
				t.Errorf("expected randon number less or eqal %d, but got %d", testCase.max, actual)
			}
		})
	}
}
