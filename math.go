/***************
* STOCK MARKET *
****************/
/*
 * Copyright (C) 2024 Carl Frank Otto III
 * All rights reserved.
 */

package main

import (
	"math"
	"math/rand"
)

func genLogRand(max float64) float64 {
	// Generate a random number between 0 and 1
	u := rand.Float64()

	// Apply logarithmic transformation and scale to the desired range
	return float64(max) * math.Log(1+u) / math.Log(volLog)
}

func calcInterest(principal float64, annualInterestRate float64) float64 {
	weeklyInterestRate := annualInterestRate / 100 / 52
	interest := principal * weeklyInterestRate
	return interest
}

func clamp(value, minVal, maxVal int) int {
	return max(min(value, maxVal), minVal)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
