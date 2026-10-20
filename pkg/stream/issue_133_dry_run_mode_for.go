package stream

import (
	"errors"
	"strings"
)

// ErrDryRunModeFor133EmptyTarget is returned when the target key is blank.
var ErrDryRunModeFor133EmptyTarget = errors.New("stream: target key cannot be empty")

// DryRunModeFor133Report summarizes deterministic diagnostics for issue #133.
type DryRunModeFor133Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateDryRunModeFor133 computes a deterministic DryRunModeFor133Report for the supplied inputs.
func EvaluateDryRunModeFor133(targetKey string, countBase, deltaUsed int) (DryRunModeFor133Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return DryRunModeFor133Report{}, ErrDryRunModeFor133EmptyTarget
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
	return DryRunModeFor133Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
