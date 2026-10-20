package stream

import "testing"

func TestEvaluateDynamicContractAddressWatchlist91(t *testing.T) {
	rep, err := EvaluateDynamicContractAddressWatchlist91("stellar-target", 40, 20)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if rep.StatusTag != "nominal" || rep.TargetKey != "stellar-target" {
		t.Fatalf("unexpected report: %+v", rep)
	}

	high, err := EvaluateDynamicContractAddressWatchlist91("stellar-target", 80, 35)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if high.StatusTag != "elevated" {
		t.Fatalf("expected elevated status, got %s", high.StatusTag)
	}

	if _, err := EvaluateDynamicContractAddressWatchlist91("   ", 1, 1); err == nil {
		t.Fatal("expected error for blank target key")
	}
}
