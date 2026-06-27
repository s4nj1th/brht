package panels

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/brht/brht/internal/engine"
	"github.com/brht/brht/internal/theme"
)

func Header(w int, s *engine.State) string {
	left := " BRHT v1.0"
	right := fmt.Sprintf("Hollywood Mode  t+%04d ", s.TickCount)
	gap := w - lipgloss.Width(left) - lipgloss.Width(right)
	if gap < 1 {
		gap = 1
	}
	line := left + strings.Repeat(" ", gap) + right
	return theme.Title.Background(theme.Background).Width(w).Render(line)
}

func AlertBar(w int, s *engine.State) string {
	if s.Alert == nil {
		return lipgloss.NewStyle().Background(theme.Background).Width(w).Render("")
	}
	style := theme.Warning
	if s.Alert.Level == "CRITICAL" {
		style = theme.Critical
	}
	text := fmt.Sprintf(" ⚠ %s: %s ", s.Alert.Level, s.Alert.Message)
	return style.Background(theme.Background).Bold(true).Width(w).Render(truncate(text, w))
}
