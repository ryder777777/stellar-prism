package stream

import (
	"errors"
	"strings"
)

// ErrLedgerHeaderScpConsensus108EmptyTarget is returned when the target key is blank.
var ErrLedgerHeaderScpConsensus108EmptyTarget = errors.New("stream: target key cannot be empty")

// LedgerHeaderScpConsensus108Report summarizes deterministic diagnostics for issue #108.
type LedgerHeaderScpConsensus108Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateLedgerHeaderScpConsensus108 computes a deterministic LedgerHeaderScpConsensus108Report for the supplied inputs.
func EvaluateLedgerHeaderScpConsensus108(targetKey string, countBase, deltaUsed int) (LedgerHeaderScpConsensus108Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return LedgerHeaderScpConsensus108Report{}, ErrLedgerHeaderScpConsensus108EmptyTarget
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
	return LedgerHeaderScpConsensus108Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
