module github.com/brht/brht

go 1.22

require (
	github.com/charmbracelet/bubbletea v1.2.4
	github.com/charmbracelet/lipgloss v1.0.0
)

require (
	github.com/aymanbagabas/go-osc52/v2 v2.0.1
	github.com/charmbracelet/x/ansi v0.4.5
	github.com/charmbracelet/x/term v0.2.1
	github.com/erikgeiser/coninput v0.0.0-20211004153227-1c3628e74d0f
	github.com/lucasb-eyer/go-colorful v1.2.0
	github.com/mattn/go-isatty v0.0.20
	github.com/mattn/go-localereader v0.0.1
	github.com/mattn/go-runewidth v0.0.16
	github.com/muesli/ansi v0.0.0-20230316100256-276c6243b2f6
	github.com/muesli/cancelreader v0.2.2
	github.com/muesli/termenv v0.15.2
	github.com/rivo/uniseg v0.4.7
	golang.org/x/sync v0.9.0
	golang.org/x/sys v0.27.0
	golang.org/x/text v0.3.8
)

replace golang.org/x/sys => github.com/golang/sys v0.27.0

replace golang.org/x/sync => github.com/golang/sync v0.9.0

replace golang.org/x/text => github.com/golang/text v0.3.8

replace golang.org/x/term => github.com/golang/term v0.26.0
