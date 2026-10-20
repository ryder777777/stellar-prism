package stream

import "testing"

func TestEvaluateTransactionSimulationPreflightStream102(t *testing.T) {
	rep, err := EvaluateTransactionSimulationPreflightStream102("stellar-target", 40, 20)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if rep.StatusTag != "nominal" || rep.TargetKey != "stellar-target" {
		t.Fatalf("unexpected report: %+v", rep)
	}

	high, err := EvaluateTransactionSimulationPreflightStream102("stellar-target", 80, 35)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if high.StatusTag != "elevated" {
		t.Fatalf("expected elevated status, got %s", high.StatusTag)
	}

	if _, err := EvaluateTransactionSimulationPreflightStream102("   ", 1, 1); err == nil {
		t.Fatal("expected error for blank target key")
	}
}
