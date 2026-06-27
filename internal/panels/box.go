package panels

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/brht/brht/internal/theme"
)

func Box(title string, outerW, outerH int, highlight bool, lines []string) string {
	innerW := outerW - 2
	innerH := outerH - 2
	if innerW < 1 {
		innerW = 1
	}
	if innerH < 1 {
		innerH = 1
	}

	header := theme.PanelTitleBar(title, innerW)
	bodyHeight := innerH - 1
	if bodyHeight < 0 {
		bodyHeight = 0
	}

	body := make([]string, 0, bodyHeight)
	for i := 0; i < bodyHeight; i++ {
		if i < len(lines) {
			body = append(body, truncate(lines[i], innerW))
		} else {
			body = append(body, "")
		}
	}

	content := header + "\n" + strings.Join(body, "\n")

	return theme.PanelBorder(highlight).
		Width(innerW).
		Height(innerH).
		Render(content)
}

func truncate(s string, w int) string {
	if w <= 0 {
		return ""
	}
	if lipgloss.Width(s) <= w {
		return s
	}
	return lipgloss.NewStyle().MaxWidth(w).Render(s)
}
