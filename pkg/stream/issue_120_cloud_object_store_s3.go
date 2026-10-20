package stream

import (
	"errors"
	"strings"
)

// ErrCloudObjectStoreS3120EmptyTarget is returned when the target key is blank.
var ErrCloudObjectStoreS3120EmptyTarget = errors.New("stream: target key cannot be empty")

// CloudObjectStoreS3120Report summarizes deterministic diagnostics for issue #120.
type CloudObjectStoreS3120Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateCloudObjectStoreS3120 computes a deterministic CloudObjectStoreS3120Report for the supplied inputs.
func EvaluateCloudObjectStoreS3120(targetKey string, countBase, deltaUsed int) (CloudObjectStoreS3120Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return CloudObjectStoreS3120Report{}, ErrCloudObjectStoreS3120EmptyTarget
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
	return CloudObjectStoreS3120Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
