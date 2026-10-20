package stream

import (
	"errors"
	"strings"
)

// ErrApacheParquetColumnarStorage114EmptyTarget is returned when the target key is blank.
var ErrApacheParquetColumnarStorage114EmptyTarget = errors.New("stream: target key cannot be empty")

// ApacheParquetColumnarStorage114Report summarizes deterministic diagnostics for issue #114.
type ApacheParquetColumnarStorage114Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateApacheParquetColumnarStorage114 computes a deterministic ApacheParquetColumnarStorage114Report for the supplied inputs.
func EvaluateApacheParquetColumnarStorage114(targetKey string, countBase, deltaUsed int) (ApacheParquetColumnarStorage114Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return ApacheParquetColumnarStorage114Report{}, ErrApacheParquetColumnarStorage114EmptyTarget
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
	return ApacheParquetColumnarStorage114Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
