package day12

import "testing"

func TestRun(t *testing.T) {
	expected := message
	if got := Run(); got != expected {
		t.Fatalf("expected %q, got %q", expected, got)
	}
}
