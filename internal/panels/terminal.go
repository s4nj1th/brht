package panels

import (
	"github.com/brht/brht/internal/engine"
	"github.com/brht/brht/internal/randutil"
)

func Terminal(w, h int, s *engine.State, glitch bool) string {
	innerH := h - 3
	if innerH < 1 {
		innerH = 1
	}
	all := s.MainLines
	start := 0
	if len(all) > innerH {
		start = len(all) - innerH
	}
	visible := all[start:]

	lines := make([]string, 0, len(visible))
	for _, l := range visible {
		if glitch && randutil.Chance(0.15) {
			l = engine.Corrupt(l)
		}
		lines = append(lines, "$ "+l)
	}
	lines = append(lines, "$ "+cursor(s.TickCount))
	return Box("MAIN TERMINAL", w, h, false, lines)
}

func cursor(tick int) string {
	if tick%2 == 0 {
		return "_"
	}
	return ""
}
