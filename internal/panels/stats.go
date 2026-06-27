package panels

import (
	"fmt"

	"github.com/brht/brht/internal/engine"
)

func Stats(w, h int, s *engine.State, glitch bool) string {
	innerH := h - 2
	if innerH < 1 {
		innerH = 1
	}
	nameW := w - 9
	if nameW < 4 {
		nameW = 4
	}

	lines := make([]string, 0, innerH)
	for i := 0; i < innerH && i < len(s.Stats); i++ {
		st := s.Stats[i]
		name := truncate(st.Name, nameW)
		line := fmt.Sprintf("%-*s %5.1f", nameW, name, st.Value)
		if glitch {
			line = engine.Corrupt(line)
		}
		lines = append(lines, line)
	}
	return Box("STATS", w, h, false, lines)
}
