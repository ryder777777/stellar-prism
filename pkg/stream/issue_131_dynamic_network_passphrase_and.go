package stream

import (
	"errors"
	"strings"
)

// ErrDynamicNetworkPassphraseAnd131EmptyTarget is returned when the target key is blank.
var ErrDynamicNetworkPassphraseAnd131EmptyTarget = errors.New("stream: target key cannot be empty")

// DynamicNetworkPassphraseAnd131Report summarizes deterministic diagnostics for issue #131.
type DynamicNetworkPassphraseAnd131Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateDynamicNetworkPassphraseAnd131 computes a deterministic DynamicNetworkPassphraseAnd131Report for the supplied inputs.
func EvaluateDynamicNetworkPassphraseAnd131(targetKey string, countBase, deltaUsed int) (DynamicNetworkPassphraseAnd131Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return DynamicNetworkPassphraseAnd131Report{}, ErrDynamicNetworkPassphraseAnd131EmptyTarget
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
	return DynamicNetworkPassphraseAnd131Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
