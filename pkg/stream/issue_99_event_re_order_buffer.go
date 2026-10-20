package stream

import (
	"errors"
	"strings"
)

// ErrEventReOrderBuffer99EmptyTarget is returned when the target key is blank.
var ErrEventReOrderBuffer99EmptyTarget = errors.New("stream: target key cannot be empty")

// EventReOrderBuffer99Report summarizes deterministic diagnostics for issue #99.
type EventReOrderBuffer99Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateEventReOrderBuffer99 computes a deterministic EventReOrderBuffer99Report for the supplied inputs.
func EvaluateEventReOrderBuffer99(targetKey string, countBase, deltaUsed int) (EventReOrderBuffer99Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return EventReOrderBuffer99Report{}, ErrEventReOrderBuffer99EmptyTarget
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
	return EventReOrderBuffer99Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
