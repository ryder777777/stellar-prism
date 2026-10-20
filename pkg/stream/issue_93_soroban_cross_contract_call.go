package stream

import (
	"errors"
	"strings"
)

// ErrSorobanCrossContractCall93EmptyTarget is returned when the target key is blank.
var ErrSorobanCrossContractCall93EmptyTarget = errors.New("stream: target key cannot be empty")

// SorobanCrossContractCall93Report summarizes deterministic diagnostics for issue #93.
type SorobanCrossContractCall93Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateSorobanCrossContractCall93 computes a deterministic SorobanCrossContractCall93Report for the supplied inputs.
func EvaluateSorobanCrossContractCall93(targetKey string, countBase, deltaUsed int) (SorobanCrossContractCall93Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return SorobanCrossContractCall93Report{}, ErrSorobanCrossContractCall93EmptyTarget
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
	return SorobanCrossContractCall93Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
