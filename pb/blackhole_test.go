package pb_test

import (
	"testing"

	"github.com/rebel-l/go-utils/pb"
)

func TestBlackHole_Finish(t *testing.T) {
	p := &pb.BlackHole{}
	p.Finish()
}

func TestBlackHole_Increment(t *testing.T) {
	p := &pb.BlackHole{}
	p.Increment()
}