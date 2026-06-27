package engine

import "testing"

func TestNewProducesSaneDefaults(t *testing.T) {
	s := New()
	if len(s.Bars) == 0 {
		t.Fatal("expected at least one progress bar")
	}
	if len(s.Stats) == 0 {
		t.Fatal("expected at least one stat")
	}
	if s.AIMessage == "" {
		t.Fatal("expected a non-empty initial AI message")
	}
}

func TestTickNeverPanics(t *testing.T) {
	s := New()
	for i := 0; i < 5000; i++ {
		s.Tick(i%7 == 0)
	}
	if s.TickCount != 5000 {
		t.Fatalf("expected TickCount to be 5000, got %d", s.TickCount)
	}
}

func TestBootSequenceCompletes(t *testing.T) {
	s := New()
	steps := 0
	for !s.StepBoot() {
		steps++
		if steps > 1000 {
			t.Fatal("boot sequence never completed")
		}
	}
	if !s.BootDone {
		t.Fatal("expected BootDone to be true after boot sequence finishes")
	}
	if len(s.BootVisible()) != len(s.BootLines) {
		t.Fatalf("expected all boot lines visible, got %d/%d", len(s.BootVisible()), len(s.BootLines))
	}
}

func TestFakeIPNeverEmpty(t *testing.T) {
	for i := 0; i < 200; i++ {
		if fakeIP() == "" {
			t.Fatal("fakeIP returned an empty string")
		}
	}
}

func TestBrainrotLineNeverEmpty(t *testing.T) {
	for i := 0; i < 200; i++ {
		if brainrotLine() == "" {
			t.Fatal("brainrotLine returned an empty string")
		}
	}
}
