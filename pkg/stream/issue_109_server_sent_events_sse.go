package stream

import (
	"errors"
	"strings"
)

// ErrServerSentEventsSse109EmptyTarget is returned when the target key is blank.
var ErrServerSentEventsSse109EmptyTarget = errors.New("stream: target key cannot be empty")

// ServerSentEventsSse109Report summarizes deterministic diagnostics for issue #109.
type ServerSentEventsSse109Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateServerSentEventsSse109 computes a deterministic ServerSentEventsSse109Report for the supplied inputs.
func EvaluateServerSentEventsSse109(targetKey string, countBase, deltaUsed int) (ServerSentEventsSse109Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return ServerSentEventsSse109Report{}, ErrServerSentEventsSse109EmptyTarget
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
	return ServerSentEventsSse109Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
