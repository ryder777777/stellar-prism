package stream

import (
	"errors"
	"strings"
)

// ErrGrpcEventStreamingService111EmptyTarget is returned when the target key is blank.
var ErrGrpcEventStreamingService111EmptyTarget = errors.New("stream: target key cannot be empty")

// GrpcEventStreamingService111Report summarizes deterministic diagnostics for issue #111.
type GrpcEventStreamingService111Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateGrpcEventStreamingService111 computes a deterministic GrpcEventStreamingService111Report for the supplied inputs.
func EvaluateGrpcEventStreamingService111(targetKey string, countBase, deltaUsed int) (GrpcEventStreamingService111Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return GrpcEventStreamingService111Report{}, ErrGrpcEventStreamingService111EmptyTarget
	}
	if countBase < 0 {
		countBase = 0
	}
	if deltaUsed < 0 {
		deltaUsed = 0
	}
	statusTag := "nominal"
	if countBase+deltaUsed > 100 {
		statusTag = "elevated"
	}
	return GrpcEventStreamingService111Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
