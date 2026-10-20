package stream

import (
	"errors"
	"strings"
)

// ErrGoogleCloudPubSub113EmptyTarget is returned when the target key is blank.
var ErrGoogleCloudPubSub113EmptyTarget = errors.New("stream: target key cannot be empty")

// GoogleCloudPubSub113Report summarizes deterministic diagnostics for issue #113.
type GoogleCloudPubSub113Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateGoogleCloudPubSub113 computes a deterministic GoogleCloudPubSub113Report for the supplied inputs.
func EvaluateGoogleCloudPubSub113(targetKey string, countBase, deltaUsed int) (GoogleCloudPubSub113Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return GoogleCloudPubSub113Report{}, ErrGoogleCloudPubSub113EmptyTarget
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
	return GoogleCloudPubSub113Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
