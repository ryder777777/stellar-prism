package stream

import (
	"errors"
	"strings"
)

// ErrElasticsearchOpensearchBulkIngestion115EmptyTarget is returned when the target key is blank.
var ErrElasticsearchOpensearchBulkIngestion115EmptyTarget = errors.New("stream: target key cannot be empty")

// ElasticsearchOpensearchBulkIngestion115Report summarizes deterministic diagnostics for issue #115.
type ElasticsearchOpensearchBulkIngestion115Report struct {
	TargetKey string
	CountBase int
	DeltaUsed int
	StatusTag string
}

// EvaluateElasticsearchOpensearchBulkIngestion115 computes a deterministic ElasticsearchOpensearchBulkIngestion115Report for the supplied inputs.
func EvaluateElasticsearchOpensearchBulkIngestion115(targetKey string, countBase, deltaUsed int) (ElasticsearchOpensearchBulkIngestion115Report, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return ElasticsearchOpensearchBulkIngestion115Report{}, ErrElasticsearchOpensearchBulkIngestion115EmptyTarget
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
	return ElasticsearchOpensearchBulkIngestion115Report{
		TargetKey: trimmed,
		CountBase: countBase,
		DeltaUsed: deltaUsed,
		StatusTag: statusTag,
	}, nil
}
