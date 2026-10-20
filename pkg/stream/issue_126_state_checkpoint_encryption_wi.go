package stream

import (
	"errors"
	"strings"
)

// ErrStateCheckpointEncryptionWith126EmptyTarget is returned when the target key is blank.
var ErrStateCheckpointEncryptionWith126EmptyTarget = errors.New("stream: target key cannot be empty")

// StateCheckpointEncryptionWith126Report summarizes deterministic diagnostics for issue #126.
type StateCheckpointEncryptionWith126Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateStateCheckpointEncryptionWith126 computes a deterministic StateCheckpointEncryptionWith126Report for the supplied inputs.
func EvaluateStateCheckpointEncryptionWith126(targetKey string, countBase, deltaUsed int) (StateCheckpointEncryptionWith126Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return StateCheckpointEncryptionWith126Report{}, ErrStateCheckpointEncryptionWith126EmptyTarget
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
	return StateCheckpointEncryptionWith126Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
