package stream

import (
	"errors"
	"strings"
)

// ErrPublicConsumerTestingUtility134EmptyTarget is returned when the target key is blank.
var ErrPublicConsumerTestingUtility134EmptyTarget = errors.New("stream: target key cannot be empty")

// PublicConsumerTestingUtility134Report summarizes deterministic diagnostics for issue #134.
type PublicConsumerTestingUtility134Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluatePublicConsumerTestingUtility134 computes a deterministic PublicConsumerTestingUtility134Report for the supplied inputs.
func EvaluatePublicConsumerTestingUtility134(targetKey string, countBase, deltaUsed int) (PublicConsumerTestingUtility134Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return PublicConsumerTestingUtility134Report{}, ErrPublicConsumerTestingUtility134EmptyTarget
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
	return PublicConsumerTestingUtility134Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
