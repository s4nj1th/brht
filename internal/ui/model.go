package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/brht/brht/internal/engine"
)

const (
	idleInterval = 120 * time.Millisecond
	bootInterval = 140 * time.Millisecond
)

type idleTickMsg struct{}
type bootTickMsg struct{}

type Model struct {
	state   *engine.State
	width   int
	height  int
	booting bool
}

func New() Model {
	return Model{
		state:   engine.New(),
		booting: true,
		width:   80,
		height:  24,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		tea.Tick(bootInterval, func(time.Time) tea.Msg { return bootTickMsg{} }),
		tea.EnterAltScreen,
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if m.booting {
			m.booting = false
			return m, tea.Tick(idleInterval, func(time.Time) tea.Msg { return idleTickMsg{} })
		}
		m.state.Tick(true)
		return m, nil

	case bootTickMsg:
		if !m.booting {
			return m, nil
		}
		done := m.state.StepBoot()
		if done {
			m.booting = false
			return m, tea.Tick(idleInterval, func(time.Time) tea.Msg { return idleTickMsg{} })
		}
		return m, tea.Tick(bootInterval, func(time.Time) tea.Msg { return bootTickMsg{} })

	case idleTickMsg:
		if m.booting {
			return m, nil
		}
		m.state.Tick(false)
		return m, tea.Tick(idleInterval, func(time.Time) tea.Msg { return idleTickMsg{} })
	}

	return m, nil
}
