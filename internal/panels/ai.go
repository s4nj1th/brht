package panels

import (
	"fmt"
	"strings"

	"github.com/brht/brht/internal/engine"
)

func AI(w, h int, s *engine.State, glitch bool) string {
	msg := s.AIMessage
	if glitch {
		msg = engine.Corrupt(msg)
	}
	barWidth := w - 14
	if barWidth < 4 {
		barWidth = 4
	}
	filled := int(float64(barWidth) * (s.AIConf / 100))
	if filled > barWidth {
		filled = barWidth
	}
	if filled < 0 {
		filled = 0
	}
	bar := strings.Repeat("█", filled) + strings.Repeat("░", barWidth-filled)

	lines := []string{
		msg,
		"",
		fmt.Sprintf("Confidence: %.0f%%", s.AIConf),
		bar,
	}
	return Box("AI", w, h, false, lines)
}
