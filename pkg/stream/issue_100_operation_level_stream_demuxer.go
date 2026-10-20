package stream

import (
	"errors"
	"strings"
)

// ErrOperationLevelStreamDemuxer100EmptyTarget is returned when the target key is blank.
var ErrOperationLevelStreamDemuxer100EmptyTarget = errors.New("stream: target key cannot be empty")

// OperationLevelStreamDemuxer100Report summarizes deterministic diagnostics for issue #100.
type OperationLevelStreamDemuxer100Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateOperationLevelStreamDemuxer100 computes a deterministic OperationLevelStreamDemuxer100Report for the supplied inputs.
func EvaluateOperationLevelStreamDemuxer100(targetKey string, countBase, deltaUsed int) (OperationLevelStreamDemuxer100Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return OperationLevelStreamDemuxer100Report{}, ErrOperationLevelStreamDemuxer100EmptyTarget
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
	return OperationLevelStreamDemuxer100Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
