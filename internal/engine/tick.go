package engine

import (
	"fmt"

	"github.com/brht/brht/internal/data"
	"github.com/brht/brht/internal/randutil"
)

func (s *State) Tick(big bool) {
	s.TickCount++

	burst := 1
	if big {
		s.KeysFed++
		burst = randutil.IntRange(2, 4)
	}

	s.updateBars(burst)
	s.updateStats(burst)
	s.updateMain(big)
	s.updateLogs(big, burst)
	s.updatePackets(big)
	s.updateAI(big)
	s.updateAlert(big)
	s.updateGlitch(big)
}

func (s *State) updateBars(burst int) {
	for i := range s.Bars {
		b := &s.Bars[i]

		switch {
		case randutil.Chance(0.01):
			b.Label = randutil.Pick(data.ProgressBarLabels)
			b.Percent = 0
			b.Vel = randutil.FloatRange(0.5, 4)
		case randutil.Chance(0.01):
			b.Percent = 100
		case randutil.Chance(0.015):
			b.Percent = randutil.FloatRange(-40, -1)
		case randutil.Chance(0.015):
			b.Percent = randutil.FloatRange(101, 260)
		case randutil.Chance(0.02):
			b.Vel = -b.Vel
		case randutil.Chance(0.008):
			b.Glitch = true
		default:
			b.Glitch = false
		}

		for n := 0; n < burst; n++ {
			b.Percent += b.Vel * randutil.FloatRange(0.3, 1.4)
		}
	}
}

func (s *State) updateStats(burst int) {
	for i := range s.Stats {
		st := &s.Stats[i]
		if randutil.Chance(0.03) {
			st.Vel = randutil.FloatRange(-4, 4)
		}
		for n := 0; n < burst; n++ {
			st.Value += st.Vel * randutil.FloatRange(0.4, 1.2)
		}
		if st.Value > 999 {
			st.Value = randutil.FloatRange(0, 50)
		}
		if st.Value < -50 {
			st.Value = randutil.FloatRange(0, 50)
		}
	}
}

func (s *State) updateMain(big bool) {
	chance := 0.10
	if big {
		chance = 0.9
	}
	if !randutil.Chance(chance) {
		return
	}
	cmd := randutil.Pick(data.MainCommands)
	dots := randutil.Pick([]string{".", "..", "...", "....", "....."})
	s.pushMain(fmt.Sprintf("%s%s", cmd, dots))
	if randutil.Chance(0.18) {
		s.pushMain(brainrotLine())
	}
}

func (s *State) updateLogs(big bool, burst int) {
	n := 0
	if big {
		n = randutil.IntRange(1, 1+burst)
	} else if randutil.Chance(0.5) {
		n = 1
	}
	for i := 0; i < n; i++ {
		level := randutil.Pick(data.LogLevels)
		var msg string
		if level == "SIGMA" || level == "AURA" || level == "BRAINROT" {
			msg = brainrotLine()
		} else {
			msg = randutil.Pick(data.LogMessages[level])
		}
		s.pushLog(level, msg)
	}
}

func (s *State) updatePackets(big bool) {
	chance := 0.12
	if big {
		chance = 0.6
	}
	if !randutil.Chance(chance) {
		return
	}
	status := randutil.Pick([]string{"OK", "OK", "OK", "DROP", "ENCRYPTED", "REROUTED"})
	s.pushPacket(Packet{
		SrcIP:    fakeIP(),
		DstIP:    fakeIP(),
		Protocol: randutil.Pick(data.NetworkProtocols),
		Status:   status,
	})
}

func (s *State) updateAI(big bool) {
	chance := 0.06
	if big {
		chance = 0.35
	}
	if randutil.Chance(chance) {
		s.AIMessage = randutil.Pick(data.AIMessages)
	}
	s.AIConf += randutil.FloatRange(-6, 6)
	if s.AIConf > 100 {
		s.AIConf = 100
	}
	if s.AIConf < 0 {
		s.AIConf = 0
	}
}

func (s *State) updateAlert(big bool) {
	if s.Alert != nil {
		s.Alert.TTL--
		if s.Alert.TTL <= 0 {
			s.Alert = nil
		}
		return
	}
	chance := 0.004
	if big {
		chance = 0.02
	}
	if randutil.Chance(chance) {
		level := "WARNING"
		msg := fmt.Sprintf("Aura dropping near %s sector.", randutil.Pick(data.NetworkProtocols))
		if randutil.Chance(0.4) {
			level = "CRITICAL"
			msg = fmt.Sprintf("%s detected. Brace for impact.", randutil.Pick(data.BrainrotFixed))
		}
		s.Alert = &Alert{Message: msg, Level: level, TTL: randutil.IntRange(8, 16)}
	}
}

func (s *State) updateGlitch(big bool) {
	if s.Glitching {
		s.GlitchTTL--
		if s.GlitchTTL <= 0 {
			s.Glitching = false
		}
		return
	}
	chance := 0.006
	if big {
		chance = 0.03
	}
	if randutil.Chance(chance) {
		s.Glitching = true
		s.GlitchTTL = randutil.IntRange(2, 5)
	}
}
