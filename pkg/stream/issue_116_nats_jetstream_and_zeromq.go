package stream

import (
	"errors"
	"strings"
)

// ErrNatsJetstreamAndZeromq116EmptyTarget is returned when the target key is blank.
var ErrNatsJetstreamAndZeromq116EmptyTarget = errors.New("stream: target key cannot be empty")

// NatsJetstreamAndZeromq116Report summarizes deterministic diagnostics for issue #116.
type NatsJetstreamAndZeromq116Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateNatsJetstreamAndZeromq116 computes a deterministic NatsJetstreamAndZeromq116Report for the supplied inputs.
func EvaluateNatsJetstreamAndZeromq116(targetKey string, countBase, deltaUsed int) (NatsJetstreamAndZeromq116Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return NatsJetstreamAndZeromq116Report{}, ErrNatsJetstreamAndZeromq116EmptyTarget
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
	return NatsJetstreamAndZeromq116Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
