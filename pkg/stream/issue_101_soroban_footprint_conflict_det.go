package stream

import (
	"errors"
	"strings"
)

// ErrSorobanFootprintConflictDetector101EmptyTarget is returned when the target key is blank.
var ErrSorobanFootprintConflictDetector101EmptyTarget = errors.New("stream: target key cannot be empty")

// SorobanFootprintConflictDetector101Report summarizes deterministic diagnostics for issue #101.
type SorobanFootprintConflictDetector101Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateSorobanFootprintConflictDetector101 computes a deterministic SorobanFootprintConflictDetector101Report for the supplied inputs.
func EvaluateSorobanFootprintConflictDetector101(targetKey string, countBase, deltaUsed int) (SorobanFootprintConflictDetector101Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return SorobanFootprintConflictDetector101Report{}, ErrSorobanFootprintConflictDetector101EmptyTarget
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
	return SorobanFootprintConflictDetector101Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
