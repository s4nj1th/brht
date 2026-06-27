package panels

import (
	"testing"

	"github.com/brht/brht/internal/engine"
)

func TestPanelsRenderAcrossSizes(t *testing.T) {
	s := engine.New()
	for i := 0; i < 50; i++ {
		s.Tick(i%3 == 0)
	}

	sizes := []struct{ w, h int }{
		{1, 1}, {5, 3}, {10, 5}, {20, 8}, {30, 10}, {80, 24}, {200, 50},
	}

	for _, sz := range sizes {
		for _, glitch := range []bool{false, true} {
			_ = System(sz.w, sz.h, s, glitch)
			_ = Terminal(sz.w, sz.h, s, glitch)
			_ = Network(sz.w, sz.h, s, glitch)
			_ = AI(sz.w, sz.h, s, glitch)
			_ = Progress(sz.w, sz.h, s, glitch)
			_ = Stats(sz.w, sz.h, s, glitch)
			_ = Logs(sz.w, sz.h, s, glitch)
			_ = Header(sz.w, s)
			_ = AlertBar(sz.w, s)
		}
	}
}
