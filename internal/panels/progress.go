package panels

import (
	"fmt"
	"strings"

	"github.com/brht/brht/internal/engine"
	"github.com/brht/brht/internal/theme"
)

func Progress(w, h int, s *engine.State, glitch bool) string {
	innerH := h - 2
	if innerH < 1 {
		innerH = 1
	}
	labelW := w/3 - 1
	if labelW < 6 {
		labelW = 6
	}
	barW := w - labelW - 9
	if barW < 4 {
		barW = 4
	}

	lines := make([]string, 0, innerH)
	for i := 0; i < innerH && i < len(s.Bars); i++ {
		b := s.Bars[i]
		pct := b.Percent
		clamped := pct
		if clamped > 100 {
			clamped = 100
		}
		if clamped < 0 {
			clamped = 0
		}
		filled := int(float64(barW) * (clamped / 100))
		if filled > barW {
			filled = barW
		}
		if filled < 0 {
			filled = 0
		}
		bar := strings.Repeat("=", filled) + strings.Repeat(" ", barW-filled)

		style := theme.Mid
		if pct > 100 || pct < 0 {
			style = theme.Accent
		}
		if b.Glitch || glitch {
			style = theme.Warning
		}

		label := truncate(b.Label, labelW)
		line := fmt.Sprintf("%-*s [%s] %4.0f%%", labelW, label, bar, pct)
		if glitch && b.Glitch {
			line = engine.Corrupt(line)
		}
		lines = append(lines, style.Render(line))
	}
	return Box("PROGRESS", w, h, false, lines)
}
