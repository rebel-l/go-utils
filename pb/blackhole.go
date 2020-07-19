package pb

import "github.com/cheggaaa/pb/v3"

// BlackHole represents a progress bar with no output.
type BlackHole struct{}

// Increment does nothing.
func (b *BlackHole) Increment() *pb.ProgressBar { return &pb.ProgressBar{} }

// Finish does nothing.
func (b *BlackHole) Finish() *pb.ProgressBar { return &pb.ProgressBar{} }
