package pb_test

import (
	"fmt"
	"testing"

	"github.com/rebel-l/go-utils/pb"
)

func TestNew(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		engine string
		total  int
	}{
		{
			engine: pb.EngineBlackhole,
			total:  123,
		},
		{
			engine: pb.EngineCheggaaa,
			total:  456,
		},
	}

	for _, testCase := range testCases {
		tc := testCase

		t.Run(tc.engine, func(t *testing.T) {
			t.Parallel()

			p := pb.New(tc.engine, tc.total)
			ty := fmt.Sprintf("%T", p)

			switch tc.engine {
			case pb.EngineBlackhole:
				if ty != "*pb.BlackHole" {
					t.Errorf("expected type BlackHole but got %s", ty)
				}
			case pb.EngineCheggaaa:
				if ty != "*pb.ProgressBar" {
					t.Errorf("expected type ProgressBar but got %s", ty)
				}
			}
		})
	}
}
