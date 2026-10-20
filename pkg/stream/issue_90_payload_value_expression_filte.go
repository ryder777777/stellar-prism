package stream

import (
	"errors"
	"strings"
)

// ErrPayloadValueExpressionFilter90EmptyTarget is returned when the target key is blank.
var ErrPayloadValueExpressionFilter90EmptyTarget = errors.New("stream: target key cannot be empty")

// PayloadValueExpressionFilter90Report summarizes deterministic diagnostics for issue #90.
type PayloadValueExpressionFilter90Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluatePayloadValueExpressionFilter90 computes a deterministic PayloadValueExpressionFilter90Report for the supplied inputs.
func EvaluatePayloadValueExpressionFilter90(targetKey string, countBase, deltaUsed int) (PayloadValueExpressionFilter90Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return PayloadValueExpressionFilter90Report{}, ErrPayloadValueExpressionFilter90EmptyTarget
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
	return PayloadValueExpressionFilter90Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
