package stream

import (
	"errors"
	"strings"
)

// ErrKubernetesHealthAndReadiness124EmptyTarget is returned when the target key is blank.
var ErrKubernetesHealthAndReadiness124EmptyTarget = errors.New("stream: target key cannot be empty")

// KubernetesHealthAndReadiness124Report summarizes deterministic diagnostics for issue #124.
type KubernetesHealthAndReadiness124Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateKubernetesHealthAndReadiness124 computes a deterministic KubernetesHealthAndReadiness124Report for the supplied inputs.
func EvaluateKubernetesHealthAndReadiness124(targetKey string, countBase, deltaUsed int) (KubernetesHealthAndReadiness124Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return KubernetesHealthAndReadiness124Report{}, ErrKubernetesHealthAndReadiness124EmptyTarget
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
	return KubernetesHealthAndReadiness124Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
