/*
 * Copyright (C) 2024 Carl Frank Otto III
 * All rights reserved.
 */

package main

import (
	"math"
	"math/rand"
	"os"
	"os/signal"
	"regexp"
	"syscall"
)

const (
	TREND_NONE = iota
	TREND_UP
	TREND_DOWN
	TREND_MAX
)

var trendSymbol [TREND_MAX]string = [TREND_MAX]string{
	"=",
	"+",
	"-",
}

func handleExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(1)
	}()
}

func (game *gameData) genLogRand(max float64) float64 {
	u := rand.Float64()

	return float64(max) * math.Log(1+u) / math.Log(game.getSettingFloat(SET_RANDLOG))
}

func NumOnly(str string) string {
	alphafilter, _ := regexp.Compile("[^0-9.]+")
	str = alphafilter.ReplaceAllString(str, "")
	return str
}

func randBool() bool {
	return rand.Float64() <= 0.5
}

func roundToCent(price float64) float64 {
	return (math.Round(price*100) / 100)
}

func floorToCent(price float64) float64 {
	return (math.Floor(price*100) / 100)
}

func roundToDollar(price float64) float64 {
	return (math.Floor(price*10000) / 10000)
}

func getTrend(a, b float64) string {
	if a > b {
		return trendSymbol[1]
	} else if b > a {
		return trendSymbol[2]
	} else {
		return trendSymbol[0]
	}
}
