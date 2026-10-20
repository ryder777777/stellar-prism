package stream

import (
	"errors"
	"strings"
)

// ErrMultiSourceEventStream97EmptyTarget is returned when the target key is blank.
var ErrMultiSourceEventStream97EmptyTarget = errors.New("stream: target key cannot be empty")

// MultiSourceEventStream97Report summarizes deterministic diagnostics for issue #97.
type MultiSourceEventStream97Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateMultiSourceEventStream97 computes a deterministic MultiSourceEventStream97Report for the supplied inputs.
func EvaluateMultiSourceEventStream97(targetKey string, countBase, deltaUsed int) (MultiSourceEventStream97Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return MultiSourceEventStream97Report{}, ErrMultiSourceEventStream97EmptyTarget
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
	return MultiSourceEventStream97Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
