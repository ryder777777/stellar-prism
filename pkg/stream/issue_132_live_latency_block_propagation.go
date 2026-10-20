package stream

import (
	"errors"
	"strings"
)

// ErrLiveLatencyBlockPropagation132EmptyTarget is returned when the target key is blank.
var ErrLiveLatencyBlockPropagation132EmptyTarget = errors.New("stream: target key cannot be empty")

// LiveLatencyBlockPropagation132Report summarizes deterministic diagnostics for issue #132.
type LiveLatencyBlockPropagation132Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateLiveLatencyBlockPropagation132 computes a deterministic LiveLatencyBlockPropagation132Report for the supplied inputs.
func EvaluateLiveLatencyBlockPropagation132(targetKey string, countBase, deltaUsed int) (LiveLatencyBlockPropagation132Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return LiveLatencyBlockPropagation132Report{}, ErrLiveLatencyBlockPropagation132EmptyTarget
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
	return LiveLatencyBlockPropagation132Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
