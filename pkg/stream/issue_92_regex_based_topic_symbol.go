package stream

import (
	"errors"
	"strings"
)

// ErrRegexBasedTopicSymbol92EmptyTarget is returned when the target key is blank.
var ErrRegexBasedTopicSymbol92EmptyTarget = errors.New("stream: target key cannot be empty")

// RegexBasedTopicSymbol92Report summarizes deterministic diagnostics for issue #92.
type RegexBasedTopicSymbol92Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateRegexBasedTopicSymbol92 computes a deterministic RegexBasedTopicSymbol92Report for the supplied inputs.
func EvaluateRegexBasedTopicSymbol92(targetKey string, countBase, deltaUsed int) (RegexBasedTopicSymbol92Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return RegexBasedTopicSymbol92Report{}, ErrRegexBasedTopicSymbol92EmptyTarget
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
	return RegexBasedTopicSymbol92Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
