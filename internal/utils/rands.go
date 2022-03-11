package utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func RandFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
