package randutils_test

import (
	"testing"

	"github.com/rebel-l/go-utils/randutils"
)

func TestInt(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		min  int
		max  int
	}{
		{
			name: "between 0 and 100",
			max:  100,
		},
		{
			name: "between 20 and 50",
			min:  20,
			max:  50,
		},
		{
			name: "between -50 and 50",
			min:  -50,
			max:  50,
		},
	}

	for _, testCase := range testCases {
		tc := testCase

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := randutils.Int(tc.min, tc.max)

			if actual < tc.min {
				t.Errorf("expected randon number greater or eqal %d, but got %d", tc.min, actual)
			}

			if actual > tc.max {
				t.Errorf("expected randon number less or eqal %d, but got %d", tc.max, actual)
			}
		})
	}
}

func TestInt_EnsureCallingTwiceReturnsDifferentNumber(t *testing.T) {
	t.Parallel()

	min := 5
	max := 10000

	var before, now int
	for i := 0; i < 10; i++ {
		now = randutils.Int(min, max)

		if i > 0 && before == now {
			t.Errorf("expected that values differ after each call")
		}

		before = now
	}
}
