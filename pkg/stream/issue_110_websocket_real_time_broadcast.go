package stream

import (
	"errors"
	"strings"
)

// ErrWebsocketRealTimeBroadcast110EmptyTarget is returned when the target key is blank.
var ErrWebsocketRealTimeBroadcast110EmptyTarget = errors.New("stream: target key cannot be empty")

// WebsocketRealTimeBroadcast110Report summarizes deterministic diagnostics for issue #110.
type WebsocketRealTimeBroadcast110Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateWebsocketRealTimeBroadcast110 computes a deterministic WebsocketRealTimeBroadcast110Report for the supplied inputs.
func EvaluateWebsocketRealTimeBroadcast110(targetKey string, countBase, deltaUsed int) (WebsocketRealTimeBroadcast110Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return WebsocketRealTimeBroadcast110Report{}, ErrWebsocketRealTimeBroadcast110EmptyTarget
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
	return WebsocketRealTimeBroadcast110Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
