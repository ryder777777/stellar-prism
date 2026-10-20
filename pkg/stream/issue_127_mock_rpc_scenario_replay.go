package stream

import (
	"errors"
	"strings"
)

// ErrMockRpcScenarioReplay127EmptyTarget is returned when the target key is blank.
var ErrMockRpcScenarioReplay127EmptyTarget = errors.New("stream: target key cannot be empty")

// MockRpcScenarioReplay127Report summarizes deterministic diagnostics for issue #127.
type MockRpcScenarioReplay127Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateMockRpcScenarioReplay127 computes a deterministic MockRpcScenarioReplay127Report for the supplied inputs.
func EvaluateMockRpcScenarioReplay127(targetKey string, countBase, deltaUsed int) (MockRpcScenarioReplay127Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return MockRpcScenarioReplay127Report{}, ErrMockRpcScenarioReplay127EmptyTarget
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
	return MockRpcScenarioReplay127Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
