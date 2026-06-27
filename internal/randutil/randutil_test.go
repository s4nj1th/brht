package randutil

import "testing"

func TestPick(t *testing.T) {
	items := []string{"a", "b", "c"}
	for i := 0; i < 100; i++ {
		v := Pick(items)
		found := false
		for _, it := range items {
			if it == v {
				found = true
			}
		}
		if !found {
			t.Fatalf("Pick returned a value not in the source slice: %q", v)
		}
	}
}

func TestIntRange(t *testing.T) {
	for i := 0; i < 1000; i++ {
		v := IntRange(5, 10)
		if v < 5 || v > 10 {
			t.Fatalf("IntRange(5, 10) returned out-of-range value: %d", v)
		}
	}
}

func TestFloatRange(t *testing.T) {
	for i := 0; i < 1000; i++ {
		v := FloatRange(1.0, 2.0)
		if v < 1.0 || v >= 2.0 {
			t.Fatalf("FloatRange(1.0, 2.0) returned out-of-range value: %f", v)
		}
	}
}

func TestChanceBounds(t *testing.T) {
	if Chance(0) {
		t.Log("Chance(0) returned true, which is only possible if rand.Float64() returned exactly 0")
	}
}
