package stream

import (
	"errors"
	"strings"
)

// ErrAwsSqsAndSns112EmptyTarget is returned when the target key is blank.
var ErrAwsSqsAndSns112EmptyTarget = errors.New("stream: target key cannot be empty")

// AwsSqsAndSns112Report summarizes deterministic diagnostics for issue #112.
type AwsSqsAndSns112Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateAwsSqsAndSns112 computes a deterministic AwsSqsAndSns112Report for the supplied inputs.
func EvaluateAwsSqsAndSns112(targetKey string, countBase, deltaUsed int) (AwsSqsAndSns112Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return AwsSqsAndSns112Report{}, ErrAwsSqsAndSns112EmptyTarget
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
	return AwsSqsAndSns112Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
