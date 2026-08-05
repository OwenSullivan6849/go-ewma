package ewma

import (
	"math"
	"testing"
)

func TestFirstSampleSeedsTheAverage(t *testing.T) {
	e := New(0.5)
	if got := e.Update(10); got != 10 {
		t.Fatalf("first Update = %v, want the sample itself (10)", got)
	}
	if got := e.Update(20); got != 15 {
		t.Fatalf("second Update = %v, want 15", got)
	}
}

func TestAlphaControlsResponsiveness(t *testing.T) {
	slow, fast := New(0.1), New(0.9)
	slow.Update(0)
	fast.Update(0)
	s, f := slow.Update(100), fast.Update(100)
	if !(f > s) {
		t.Errorf("higher alpha should react faster: fast=%v slow=%v", f, s)
	}
	if math.Abs(f-90) > 1e-9 {
		t.Errorf("fast = %v, want 90", f)
	}
}
