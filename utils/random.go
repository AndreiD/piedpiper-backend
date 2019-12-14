package utils

import (
	"math/rand"
	"time"
)

// RandFloat .
func RandFloat(min, max int) float64 {
	rand.Seed(time.Now().UTC().UnixNano())

	randNr := RandInts(0, 10)
	if randNr < 5 {
		// return a int number
		return float64(RandInts(min, max))
	}

	if randNr < 8 {
		return float64((RandInts(min, max) / 100) * 100)
	}
	// else

	expfloat := rand.ExpFloat64()
	if expfloat > float64(max) {
		return float64(max)
	}

	value := float64(min) + expfloat*100
	if value > float64(max) {
		return float64(max)
	}
	if value < float64(min) {
		return float64(min)
	}
	return value
}

// RandFloat .
func RandFloatf(min, max float64) float64 {
	rand.Seed(time.Now().UTC().UnixNano())
	value := min + rand.Float64()*(max-min)
	return value
}

// RandInts .
func RandInts(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min+1)
}
