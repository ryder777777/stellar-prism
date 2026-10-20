package stream

import (
	"errors"
	"strings"
)

// ErrStructuredAuditTrailCompliance135EmptyTarget is returned when the target key is blank.
var ErrStructuredAuditTrailCompliance135EmptyTarget = errors.New("stream: target key cannot be empty")

// StructuredAuditTrailCompliance135Report summarizes deterministic diagnostics for issue #135.
type StructuredAuditTrailCompliance135Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateStructuredAuditTrailCompliance135 computes a deterministic StructuredAuditTrailCompliance135Report for the supplied inputs.
func EvaluateStructuredAuditTrailCompliance135(targetKey string, countBase, deltaUsed int) (StructuredAuditTrailCompliance135Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return StructuredAuditTrailCompliance135Report{}, ErrStructuredAuditTrailCompliance135EmptyTarget
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
	return StructuredAuditTrailCompliance135Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
