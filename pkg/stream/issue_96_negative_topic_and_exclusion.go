package stream

import (
	"errors"
	"strings"
)

// ErrNegativeTopicAndExclusion96EmptyTarget is returned when the target key is blank.
var ErrNegativeTopicAndExclusion96EmptyTarget = errors.New("stream: target key cannot be empty")

// NegativeTopicAndExclusion96Report summarizes deterministic diagnostics for issue #96.
type NegativeTopicAndExclusion96Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateNegativeTopicAndExclusion96 computes a deterministic NegativeTopicAndExclusion96Report for the supplied inputs.
func EvaluateNegativeTopicAndExclusion96(targetKey string, countBase, deltaUsed int) (NegativeTopicAndExclusion96Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return NegativeTopicAndExclusion96Report{}, ErrNegativeTopicAndExclusion96EmptyTarget
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
	return NegativeTopicAndExclusion96Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
