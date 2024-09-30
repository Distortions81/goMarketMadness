/*
 * Copyright (C) 2024 Carl Frank Otto III
 * All rights reserved.
 */

package main

import (
	"math"
	"math/rand"
)

func (game *gameData) genLogRand(max float64) float64 {
	u := rand.Float64()

	return float64(max) * math.Log(1+u) / math.Log(game.getSettingFloat(SET_RANDLOG))
}

func clampInt(value, minVal, maxVal int) int {
	return maxInt(minInt(value, maxVal), minVal)
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
