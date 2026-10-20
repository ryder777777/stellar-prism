package stream

import (
	"errors"
	"strings"
)

// ErrDynamicContractAddressWatchlist91EmptyTarget is returned when the target key is blank.
var ErrDynamicContractAddressWatchlist91EmptyTarget = errors.New("stream: target key cannot be empty")

// DynamicContractAddressWatchlist91Report summarizes deterministic diagnostics for issue #91.
type DynamicContractAddressWatchlist91Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateDynamicContractAddressWatchlist91 computes a deterministic DynamicContractAddressWatchlist91Report for the supplied inputs.
func EvaluateDynamicContractAddressWatchlist91(targetKey string, countBase, deltaUsed int) (DynamicContractAddressWatchlist91Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return DynamicContractAddressWatchlist91Report{}, ErrDynamicContractAddressWatchlist91EmptyTarget
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
	return DynamicContractAddressWatchlist91Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
