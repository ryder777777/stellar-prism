package stream

import (
	"errors"
	"strings"
)

// ErrBidirectionalStreamRewindReplay121EmptyTarget is returned when the target key is blank.
var ErrBidirectionalStreamRewindReplay121EmptyTarget = errors.New("stream: target key cannot be empty")

// BidirectionalStreamRewindReplay121Report summarizes deterministic diagnostics for issue #121.
type BidirectionalStreamRewindReplay121Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateBidirectionalStreamRewindReplay121 computes a deterministic BidirectionalStreamRewindReplay121Report for the supplied inputs.
func EvaluateBidirectionalStreamRewindReplay121(targetKey string, countBase, deltaUsed int) (BidirectionalStreamRewindReplay121Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return BidirectionalStreamRewindReplay121Report{}, ErrBidirectionalStreamRewindReplay121EmptyTarget
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
	return BidirectionalStreamRewindReplay121Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
