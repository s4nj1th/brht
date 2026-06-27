package panels

import (
	"fmt"

	"github.com/brht/brht/internal/engine"
	"github.com/brht/brht/internal/theme"
)

func Logs(w, h int, s *engine.State, glitch bool) string {
	innerH := h - 2
	if innerH < 1 {
		innerH = 1
	}
	all := s.Logs
	start := 0
	if len(all) > innerH {
		start = len(all) - innerH
	}
	visible := all[start:]

	lines := make([]string, 0, len(visible))
	for _, e := range visible {
		style := theme.LevelStyle(e.Level)
		msg := e.Message
		if glitch {
			msg = engine.Corrupt(msg)
		}
		line := fmt.Sprintf("%s [%-8s] %s", e.Stamp, e.Level, msg)
		lines = append(lines, style.Render(line))
	}
	title := "LOGS"
	if len(s.Logs) > innerH {
		title = "LOGS ↑"
	}
	return Box(title, w, h, false, lines)
}
