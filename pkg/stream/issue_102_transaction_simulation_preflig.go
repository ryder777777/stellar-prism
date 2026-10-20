package stream

import (
	"errors"
	"strings"
)

// ErrTransactionSimulationPreflightStream102EmptyTarget is returned when the target key is blank.
var ErrTransactionSimulationPreflightStream102EmptyTarget = errors.New("stream: target key cannot be empty")

// TransactionSimulationPreflightStream102Report summarizes deterministic diagnostics for issue #102.
type TransactionSimulationPreflightStream102Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateTransactionSimulationPreflightStream102 computes a deterministic TransactionSimulationPreflightStream102Report for the supplied inputs.
func EvaluateTransactionSimulationPreflightStream102(targetKey string, countBase, deltaUsed int) (TransactionSimulationPreflightStream102Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return TransactionSimulationPreflightStream102Report{}, ErrTransactionSimulationPreflightStream102EmptyTarget
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
	return TransactionSimulationPreflightStream102Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
