package stream

import (
	"errors"
	"strings"
)

// ErrSmartContractUpgradeState105EmptyTarget is returned when the target key is blank.
var ErrSmartContractUpgradeState105EmptyTarget = errors.New("stream: target key cannot be empty")

// SmartContractUpgradeState105Report summarizes deterministic diagnostics for issue #105.
type SmartContractUpgradeState105Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateSmartContractUpgradeState105 computes a deterministic SmartContractUpgradeState105Report for the supplied inputs.
func EvaluateSmartContractUpgradeState105(targetKey string, countBase, deltaUsed int) (SmartContractUpgradeState105Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return SmartContractUpgradeState105Report{}, ErrSmartContractUpgradeState105EmptyTarget
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
	return SmartContractUpgradeState105Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
