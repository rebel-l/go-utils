package pb_test

import (
	"fmt"
	"testing"

	"github.com/rebel-l/go-utils/pb"
)

func TestNew(t *testing.T) {
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
		t.Run(testCase.engine, func(t *testing.T) {
			p := pb.New(testCase.engine, testCase.total)
			ty := fmt.Sprintf("%T", p)

			switch testCase.engine {
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
