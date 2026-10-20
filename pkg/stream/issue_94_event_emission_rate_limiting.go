package stream

import (
	"errors"
	"strings"
)

// ErrEventEmissionRateLimiting94EmptyTarget is returned when the target key is blank.
var ErrEventEmissionRateLimiting94EmptyTarget = errors.New("stream: target key cannot be empty")

// EventEmissionRateLimiting94Report summarizes deterministic diagnostics for issue #94.
type EventEmissionRateLimiting94Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateEventEmissionRateLimiting94 computes a deterministic EventEmissionRateLimiting94Report for the supplied inputs.
func EvaluateEventEmissionRateLimiting94(targetKey string, countBase, deltaUsed int) (EventEmissionRateLimiting94Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return EventEmissionRateLimiting94Report{}, ErrEventEmissionRateLimiting94EmptyTarget
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
	return EventEmissionRateLimiting94Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
