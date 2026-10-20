package stream

import (
	"errors"
	"strings"
)

// ErrWriteAheadLogWal119EmptyTarget is returned when the target key is blank.
var ErrWriteAheadLogWal119EmptyTarget = errors.New("stream: target key cannot be empty")

// WriteAheadLogWal119Report summarizes deterministic diagnostics for issue #119.
type WriteAheadLogWal119Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateWriteAheadLogWal119 computes a deterministic WriteAheadLogWal119Report for the supplied inputs.
func EvaluateWriteAheadLogWal119(targetKey string, countBase, deltaUsed int) (WriteAheadLogWal119Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return WriteAheadLogWal119Report{}, ErrWriteAheadLogWal119EmptyTarget
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
	return WriteAheadLogWal119Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
