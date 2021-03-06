package memdev

import "testing"

func TestInfo(t *testing.T) {
	memInfo, err := Info()
	if err != nil {
		t.Fatalf("Expected nil error but got %v", err)
	}
	if len(memInfo) < 0 {
		t.Fatalf("Expected Memory >= 0, but got %d", len(memInfo))
	}
}
