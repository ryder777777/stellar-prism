package stream

import (
	"errors"
	"strings"
)

// ErrInteractivePayloadInspectorAnd128EmptyTarget is returned when the target key is blank.
var ErrInteractivePayloadInspectorAnd128EmptyTarget = errors.New("stream: target key cannot be empty")

// InteractivePayloadInspectorAnd128Report summarizes deterministic diagnostics for issue #128.
type InteractivePayloadInspectorAnd128Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateInteractivePayloadInspectorAnd128 computes a deterministic InteractivePayloadInspectorAnd128Report for the supplied inputs.
func EvaluateInteractivePayloadInspectorAnd128(targetKey string, countBase, deltaUsed int) (InteractivePayloadInspectorAnd128Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return InteractivePayloadInspectorAnd128Report{}, ErrInteractivePayloadInspectorAnd128EmptyTarget
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
	return InteractivePayloadInspectorAnd128Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
