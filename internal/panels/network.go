package panels

import (
	"fmt"

	"github.com/brht/brht/internal/engine"
	"github.com/brht/brht/internal/theme"
)

func Network(w, h int, s *engine.State, glitch bool) string {
	innerH := h - 2
	if innerH < 1 {
		innerH = 1
	}
	all := s.Packets
	start := 0
	if len(all) > innerH {
		start = len(all) - innerH
	}
	visible := all[start:]

	lines := make([]string, 0, len(visible))
	for _, p := range visible {
		style := theme.Mid
		switch p.Status {
		case "DROP":
			style = theme.Warning
		case "ENCRYPTED":
			style = theme.Accent
		}
		line := fmt.Sprintf("%s>%s [%s] %s", p.SrcIP, p.DstIP, p.Protocol, p.Status)
		if glitch {
			line = engine.Corrupt(line)
		}
		lines = append(lines, style.Render(line))
	}
	return Box("NETWORK", w, h, false, lines)
}
