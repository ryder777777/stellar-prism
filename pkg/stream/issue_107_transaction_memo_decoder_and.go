package stream

import (
	"errors"
	"strings"
)

// ErrTransactionMemoDecoderAnd107EmptyTarget is returned when the target key is blank.
var ErrTransactionMemoDecoderAnd107EmptyTarget = errors.New("stream: target key cannot be empty")

// TransactionMemoDecoderAnd107Report summarizes deterministic diagnostics for issue #107.
type TransactionMemoDecoderAnd107Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateTransactionMemoDecoderAnd107 computes a deterministic TransactionMemoDecoderAnd107Report for the supplied inputs.
func EvaluateTransactionMemoDecoderAnd107(targetKey string, countBase, deltaUsed int) (TransactionMemoDecoderAnd107Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return TransactionMemoDecoderAnd107Report{}, ErrTransactionMemoDecoderAnd107EmptyTarget
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
	return TransactionMemoDecoderAnd107Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
