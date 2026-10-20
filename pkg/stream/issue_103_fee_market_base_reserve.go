package stream

import (
	"errors"
	"strings"
)

// ErrFeeMarketBaseReserve103EmptyTarget is returned when the target key is blank.
var ErrFeeMarketBaseReserve103EmptyTarget = errors.New("stream: target key cannot be empty")

// FeeMarketBaseReserve103Report summarizes deterministic diagnostics for issue #103.
type FeeMarketBaseReserve103Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateFeeMarketBaseReserve103 computes a deterministic FeeMarketBaseReserve103Report for the supplied inputs.
func EvaluateFeeMarketBaseReserve103(targetKey string, countBase, deltaUsed int) (FeeMarketBaseReserve103Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return FeeMarketBaseReserve103Report{}, ErrFeeMarketBaseReserve103EmptyTarget
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
	return FeeMarketBaseReserve103Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
