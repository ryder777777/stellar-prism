package stream

import (
	"errors"
	"strings"
)

// ErrGracefulDrainWithConfigurable125EmptyTarget is returned when the target key is blank.
var ErrGracefulDrainWithConfigurable125EmptyTarget = errors.New("stream: target key cannot be empty")

// GracefulDrainWithConfigurable125Report summarizes deterministic diagnostics for issue #125.
type GracefulDrainWithConfigurable125Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateGracefulDrainWithConfigurable125 computes a deterministic GracefulDrainWithConfigurable125Report for the supplied inputs.
func EvaluateGracefulDrainWithConfigurable125(targetKey string, countBase, deltaUsed int) (GracefulDrainWithConfigurable125Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return GracefulDrainWithConfigurable125Report{}, ErrGracefulDrainWithConfigurable125EmptyTarget
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
	return GracefulDrainWithConfigurable125Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
