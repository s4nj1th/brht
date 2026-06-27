package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/brht/brht/internal/panels"
	"github.com/brht/brht/internal/theme"
)

func (m Model) View() string {
	if m.booting {
		return m.viewBoot()
	}
	return m.viewMain()
}

func (m Model) viewBoot() string {
	lines := m.state.BootVisible()
	var b strings.Builder
	pad := m.height/2 - len(lines)/2
	if pad < 0 {
		pad = 1
	}
	b.WriteString(strings.Repeat("\n", pad))
	for _, l := range lines {
		b.WriteString(theme.Base.Render("  > " + l))
		b.WriteString("\n")
	}
	body := lipgloss.NewStyle().
		Background(theme.Background).
		Width(m.width).
		Height(m.height).
		Render(b.String())
	return body
}

func (m Model) viewMain() string {
	w, h := m.width, m.height
	if w < 20 {
		w = 20
	}
	if h < 12 {
		h = 12
	}
	glitch := m.state.Glitching

	header := panels.Header(w, m.state)
	alert := panels.AlertBar(w, m.state)

	logsH := h / 3
	if logsH < 6 {
		logsH = 6
	}
	gridH := h - 2 - logsH
	if gridH < 6 {
		gridH = 6
	}
	rowH := gridH / 2

	col1 := w / 3
	col3 := w / 3
	col2 := w - col1 - col3

	row1 := lipgloss.JoinHorizontal(lipgloss.Top,
		panels.System(col1, rowH, m.state, glitch),
		panels.Terminal(col2, rowH, m.state, glitch),
		panels.Network(col3, rowH, m.state, glitch),
	)
	row2 := lipgloss.JoinHorizontal(lipgloss.Top,
		panels.AI(col1, rowH, m.state, glitch),
		panels.Progress(col2, rowH, m.state, glitch),
		panels.Stats(col3, rowH, m.state, glitch),
	)
	logs := panels.Logs(w, logsH, m.state, glitch)

	full := lipgloss.JoinVertical(lipgloss.Left, header, alert, row1, row2, logs)

	return lipgloss.NewStyle().Background(theme.Background).Render(full)
}
