package engine

import (
	"strings"

	"github.com/brht/brht/internal/data"
	"github.com/brht/brht/internal/randutil"
)

func Corrupt(s string) string {
	if s == "" {
		return s
	}
	r := []rune(s)
	hits := randutil.IntRange(1, 3)
	for i := 0; i < hits; i++ {
		pos := randutil.IntRange(0, len(r)-1)
		if r[pos] == ' ' {
			continue
		}
		if randutil.Chance(0.5) {
			r[pos] = randutil.Pick(data.GlitchGlyphs)
		} else {
			dup := string(r[pos])
			return string(r[:pos]) + dup + string(r[pos:])
		}
	}
	return string(r)
}

func CorruptCursorJump(s string) string {
	noise := strings.Repeat(string(randutil.Pick(data.GlitchGlyphs)), randutil.IntRange(1, 3))
	return noise + s
}
