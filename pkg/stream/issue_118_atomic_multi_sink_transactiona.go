package stream

import (
	"errors"
	"strings"
)

// ErrAtomicMultiSinkTransactional118EmptyTarget is returned when the target key is blank.
var ErrAtomicMultiSinkTransactional118EmptyTarget = errors.New("stream: target key cannot be empty")

// AtomicMultiSinkTransactional118Report summarizes deterministic diagnostics for issue #118.
type AtomicMultiSinkTransactional118Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateAtomicMultiSinkTransactional118 computes a deterministic AtomicMultiSinkTransactional118Report for the supplied inputs.
func EvaluateAtomicMultiSinkTransactional118(targetKey string, countBase, deltaUsed int) (AtomicMultiSinkTransactional118Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return AtomicMultiSinkTransactional118Report{}, ErrAtomicMultiSinkTransactional118EmptyTarget
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
	return AtomicMultiSinkTransactional118Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
