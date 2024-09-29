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

func genLogRand(game *gameData, max float64) float64 {
	u := rand.Float64()

	return float64(max) * math.Log(1+u) / math.Log(game.settings.randLogarithm)
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
