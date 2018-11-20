package random

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// Rand01 returns, as a float64, a pseudo-random number in [0.0,1.0).
func Rand01() float64 {
	return rand.Float64()
}

// Rand0n returns, as an int64, a non-negative pseudo-random number in [0,n).
func Rand0n(n int64) int64 {
	return rand.Int63n(n)
}

// HitJudge hit judge func with rate k, k should in [0.0, 1.0]
func HitJudge(k float64) bool {
	if k <= 0.0 {
		return false
	}
	if k >= 1.0 {
		return true
	}
	return Rand01() <= k
}
