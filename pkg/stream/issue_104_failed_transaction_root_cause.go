package stream

import (
	"errors"
	"strings"
)

// ErrFailedTransactionRootCause104EmptyTarget is returned when the target key is blank.
var ErrFailedTransactionRootCause104EmptyTarget = errors.New("stream: target key cannot be empty")

// FailedTransactionRootCause104Report summarizes deterministic diagnostics for issue #104.
type FailedTransactionRootCause104Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateFailedTransactionRootCause104 computes a deterministic FailedTransactionRootCause104Report for the supplied inputs.
func EvaluateFailedTransactionRootCause104(targetKey string, countBase, deltaUsed int) (FailedTransactionRootCause104Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return FailedTransactionRootCause104Report{}, ErrFailedTransactionRootCause104EmptyTarget
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
	return FailedTransactionRootCause104Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
