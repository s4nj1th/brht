package theme

import "github.com/charmbracelet/lipgloss"

var (
	Background  = lipgloss.Color("#020503")
	GreenBright = lipgloss.Color("#39FF14")
	GreenMid    = lipgloss.Color("#22C55E")
	GreenDim    = lipgloss.Color("#0F5C2E")
	GreenFaint  = lipgloss.Color("#073B1D")
	Red         = lipgloss.Color("#FF3B3B")
	Yellow      = lipgloss.Color("#E8D44D")
	Cyan        = lipgloss.Color("#3FD9C7")
	White       = lipgloss.Color("#EAFBEF")
)

var Base = lipgloss.NewStyle().Foreground(GreenBright)

var Dim = lipgloss.NewStyle().Foreground(GreenDim)

var Mid = lipgloss.NewStyle().Foreground(GreenMid)

var Title = lipgloss.NewStyle().Foreground(GreenBright).Bold(true)

var (
	Critical  = lipgloss.NewStyle().Foreground(Red).Bold(true)
	Warning   = lipgloss.NewStyle().Foreground(Yellow)
	Accent    = lipgloss.NewStyle().Foreground(Cyan)
	RareWhite = lipgloss.NewStyle().Foreground(White).Bold(true)
)

func LevelStyle(level string) lipgloss.Style {
	switch level {
	case "CRITICAL":
		return Critical
	case "ERROR":
		return lipgloss.NewStyle().Foreground(Red)
	case "WARNING":
		return Warning
	case "SIGMA", "AURA", "BRAINROT":
		return Accent
	default:
		return Mid
	}
}

func PanelBorder(highlight bool) lipgloss.Style {
	c := GreenDim
	if highlight {
		c = GreenBright
	}
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(c).
		Background(Background)
}

func PanelTitleBar(label string, width int) string {
	return Title.Background(Background).Width(width).Render(" " + label + " ")
}
