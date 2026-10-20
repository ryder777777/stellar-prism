package stream

import (
	"errors"
	"strings"
)

// ErrStatefulEventAggregatorAnd95EmptyTarget is returned when the target key is blank.
var ErrStatefulEventAggregatorAnd95EmptyTarget = errors.New("stream: target key cannot be empty")

// StatefulEventAggregatorAnd95Report summarizes deterministic diagnostics for issue #95.
type StatefulEventAggregatorAnd95Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateStatefulEventAggregatorAnd95 computes a deterministic StatefulEventAggregatorAnd95Report for the supplied inputs.
func EvaluateStatefulEventAggregatorAnd95(targetKey string, countBase, deltaUsed int) (StatefulEventAggregatorAnd95Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return StatefulEventAggregatorAnd95Report{}, ErrStatefulEventAggregatorAnd95EmptyTarget
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
	return StatefulEventAggregatorAnd95Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
