package stream

import (
	"errors"
	"strings"
)

// ErrContractEventCodeGenerator129EmptyTarget is returned when the target key is blank.
var ErrContractEventCodeGenerator129EmptyTarget = errors.New("stream: target key cannot be empty")

// ContractEventCodeGenerator129Report summarizes deterministic diagnostics for issue #129.
type ContractEventCodeGenerator129Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateContractEventCodeGenerator129 computes a deterministic ContractEventCodeGenerator129Report for the supplied inputs.
func EvaluateContractEventCodeGenerator129(targetKey string, countBase, deltaUsed int) (ContractEventCodeGenerator129Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return ContractEventCodeGenerator129Report{}, ErrContractEventCodeGenerator129EmptyTarget
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
	return ContractEventCodeGenerator129Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
