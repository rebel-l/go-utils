package pb

import "github.com/cheggaaa/pb/v3"

const (
	// EngineBlackhole represents the string for engine BlackHole
	EngineBlackhole = "blackhole"

	// EngineCheggaaa represents the string for engine Cheggaaa
	EngineCheggaaa = "cheggaaa"
)

// Progressor provides methods to steer a progress bar.
type Progressor interface {
	Increment() *pb.ProgressBar
	Finish() *pb.ProgressBar
}

// New returns a new progress bar initialized with the total number of elements.
func New(engine string, total int) Progressor {
	var progressBar Progressor

	switch engine {
	case EngineBlackhole:
		progressBar = &BlackHole{}
	case EngineCheggaaa:
		progressBar = pb.New(total)
	}

	return progressBar
}
