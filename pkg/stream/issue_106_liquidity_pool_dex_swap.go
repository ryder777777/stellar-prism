package stream

import (
	"errors"
	"strings"
)

// ErrLiquidityPoolDexSwap106EmptyTarget is returned when the target key is blank.
var ErrLiquidityPoolDexSwap106EmptyTarget = errors.New("stream: target key cannot be empty")

// LiquidityPoolDexSwap106Report summarizes deterministic diagnostics for issue #106.
type LiquidityPoolDexSwap106Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateLiquidityPoolDexSwap106 computes a deterministic LiquidityPoolDexSwap106Report for the supplied inputs.
func EvaluateLiquidityPoolDexSwap106(targetKey string, countBase, deltaUsed int) (LiquidityPoolDexSwap106Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return LiquidityPoolDexSwap106Report{}, ErrLiquidityPoolDexSwap106EmptyTarget
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
	return LiquidityPoolDexSwap106Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
