package stream

import (
	"errors"
	"strings"
)

// ErrDeadLetterQueueDlq122EmptyTarget is returned when the target key is blank.
var ErrDeadLetterQueueDlq122EmptyTarget = errors.New("stream: target key cannot be empty")

// DeadLetterQueueDlq122Report summarizes deterministic diagnostics for issue #122.
type DeadLetterQueueDlq122Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateDeadLetterQueueDlq122 computes a deterministic DeadLetterQueueDlq122Report for the supplied inputs.
func EvaluateDeadLetterQueueDlq122(targetKey string, countBase, deltaUsed int) (DeadLetterQueueDlq122Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return DeadLetterQueueDlq122Report{}, ErrDeadLetterQueueDlq122EmptyTarget
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
	return DeadLetterQueueDlq122Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
