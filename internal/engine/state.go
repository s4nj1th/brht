package engine

import (
	"fmt"
	"time"

	"github.com/brht/brht/internal/data"
	"github.com/brht/brht/internal/randutil"
)

type LogEntry struct {
	Level   string
	Message string
	Stamp   string
}

type ProgressBar struct {
	Label   string
	Percent float64
	Vel     float64
	Glitch  bool
}

type Stat struct {
	Name  string
	Value float64
	Vel   float64
}

type Packet struct {
	SrcIP    string
	DstIP    string
	Protocol string
	Status   string
}

type Alert struct {
	Message string
	Level   string
	TTL     int
}

type State struct {
	Logs      []LogEntry
	Bars      []ProgressBar
	Stats     []Stat
	Packets   []Packet
	MainLines []string
	AIMessage string
	AIConf    float64
	Alert     *Alert
	Glitching bool
	GlitchTTL int
	KeysFed   int
	TickCount int
	BootDone  bool
	BootLines []string
	bootIndex int
}

const (
	maxLogs     = 300
	maxMain     = 200
	maxPackets  = 30
	numBars     = 11
	visibleLogs = 14
)

func New() *State {
	s := &State{
		BootLines: data.BootSequence,
	}
	s.Bars = make([]ProgressBar, numBars)
	for i := range s.Bars {
		s.Bars[i] = ProgressBar{
			Label:   randutil.Pick(data.ProgressBarLabels),
			Percent: randutil.FloatRange(0, 60),
			Vel:     randutil.FloatRange(0.5, 4),
		}
	}
	s.Stats = make([]Stat, len(data.StatNames))
	for i, name := range data.StatNames {
		s.Stats[i] = Stat{
			Name:  name,
			Value: randutil.FloatRange(10, 90),
			Vel:   randutil.FloatRange(-3, 3),
		}
	}
	s.AIMessage = randutil.Pick(data.AIMessages)
	s.AIConf = randutil.FloatRange(20, 80)
	return s
}

func (s *State) StepBoot() bool {
	if s.bootIndex >= len(s.BootLines) {
		s.BootDone = true
		return true
	}
	s.bootIndex++
	if s.bootIndex >= len(s.BootLines) {
		s.BootDone = true
	}
	return s.BootDone
}

func (s *State) BootVisible() []string {
	if s.bootIndex > len(s.BootLines) {
		return s.BootLines
	}
	return s.BootLines[:s.bootIndex]
}

func (s *State) pushLog(level, msg string) {
	s.Logs = append(s.Logs, LogEntry{Level: level, Message: msg, Stamp: stamp()})
	if len(s.Logs) > maxLogs {
		s.Logs = s.Logs[len(s.Logs)-maxLogs:]
	}
}

func (s *State) pushMain(line string) {
	s.MainLines = append(s.MainLines, line)
	if len(s.MainLines) > maxMain {
		s.MainLines = s.MainLines[len(s.MainLines)-maxMain:]
	}
}

func (s *State) pushPacket(p Packet) {
	s.Packets = append(s.Packets, p)
	if len(s.Packets) > maxPackets {
		s.Packets = s.Packets[len(s.Packets)-maxPackets:]
	}
}

func stamp() string {
	return time.Now().Format("15:04:05")
}

func (s *State) VisibleLogs() []LogEntry {
	if len(s.Logs) <= visibleLogs {
		return s.Logs
	}
	return s.Logs[len(s.Logs)-visibleLogs:]
}

func fakeIP() string {
	switch randutil.IntRange(0, 3) {
	case 0:
		return fmt.Sprintf("%d.%d.%d.%d", randutil.IntRange(100, 999), randutil.IntRange(0, 999), randutil.IntRange(0, 999), randutil.IntRange(0, 999))
	case 1:
		return fmt.Sprintf("10.0.0.%s", randutil.Pick([]string{"brain", "aura", "sigma", "ohio", "rizz"}))
	case 2:
		return "420.420.420.420"
	default:
		return fmt.Sprintf("%d.%d.%d.%d", randutil.IntRange(1, 223), randutil.IntRange(0, 255), randutil.IntRange(0, 255), randutil.IntRange(0, 255))
	}
}

func brainrotLine() string {
	if randutil.Chance(0.35) {
		return randutil.Pick(data.BrainrotFixed)
	}
	return fmt.Sprintf("%s %s", randutil.Pick(data.BrainrotSubjects), randutil.Pick(data.BrainrotActions))
}
