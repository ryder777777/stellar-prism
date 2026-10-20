package stream

import (
	"errors"
	"strings"
)

// ErrCompactCsvAndTsv117EmptyTarget is returned when the target key is blank.
var ErrCompactCsvAndTsv117EmptyTarget = errors.New("stream: target key cannot be empty")

// CompactCsvAndTsv117Report summarizes deterministic diagnostics for issue #117.
type CompactCsvAndTsv117Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateCompactCsvAndTsv117 computes a deterministic CompactCsvAndTsv117Report for the supplied inputs.
func EvaluateCompactCsvAndTsv117(targetKey string, countBase, deltaUsed int) (CompactCsvAndTsv117Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return CompactCsvAndTsv117Report{}, ErrCompactCsvAndTsv117EmptyTarget
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
	return CompactCsvAndTsv117Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
