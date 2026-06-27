package panels

import (
	"fmt"

	"github.com/brht/brht/internal/engine"
	"github.com/brht/brht/internal/randutil"
)

var hostnames = []string{"GIBSON-7", "MAINFRAME-OHIO", "SIGMA-CORE", "NODE-67", "AURA-PRIME"}

var bootedHostname = randutil.Pick(hostnames)

func System(w, h int, s *engine.State, glitch bool) string {
	lines := []string{
		fmt.Sprintf("Host: %s", bootedHostname),
		"OS: Linux (Fake) 6.9.0-sigma",
		"Mode: HOLLYWOOD",
		fmt.Sprintf("Uptime: %dt", s.TickCount),
		fmt.Sprintf("Keys fed: %d", s.KeysFed),
		fmt.Sprintf("Threat lvl: %d%%", randutil.IntRange(0, 100)),
		fmt.Sprintf("Aura core: %.1f°", randutil.FloatRange(60, 99)),
		"Status: definitely fine",
	}
	if glitch {
		for i := range lines {
			if randutil.Chance(0.3) {
				lines[i] = engine.Corrupt(lines[i])
			}
		}
	}
	return Box("SYSTEM", w, h, false, lines)
}
