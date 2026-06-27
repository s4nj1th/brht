package randutil

import "math/rand"

func Pick[T any](items []T) T {
	return items[rand.Intn(len(items))]
}

func IntRange(min, max int) int {
	if max <= min {
		return min
	}
	return min + rand.Intn(max-min+1)
}

func FloatRange(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func Chance(p float64) bool {
	return rand.Float64() < p
}
