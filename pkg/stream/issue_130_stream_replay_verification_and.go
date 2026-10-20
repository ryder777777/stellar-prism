package stream

import (
	"errors"
	"strings"
)

// ErrStreamReplayVerificationAnd130EmptyTarget is returned when the target key is blank.
var ErrStreamReplayVerificationAnd130EmptyTarget = errors.New("stream: target key cannot be empty")

// StreamReplayVerificationAnd130Report summarizes deterministic diagnostics for issue #130.
type StreamReplayVerificationAnd130Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateStreamReplayVerificationAnd130 computes a deterministic StreamReplayVerificationAnd130Report for the supplied inputs.
func EvaluateStreamReplayVerificationAnd130(targetKey string, countBase, deltaUsed int) (StreamReplayVerificationAnd130Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return StreamReplayVerificationAnd130Report{}, ErrStreamReplayVerificationAnd130EmptyTarget
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
	return StreamReplayVerificationAnd130Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
