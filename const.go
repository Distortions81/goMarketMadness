/***************
* STOCK MARKET *
****************/
/*
 * Copyright (C) 2024 Carl Frank Otto III
 * All rights reserved.
 */

package main

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

var defSettings = settingsData{
	maxPlayers: 25, maxWeeks: 52, maxNameLen: 64, randLogarithm: 100,
	maxLoanCount: 10, maxShares: 10000,

	startingMoney: 5000, minSigma: 3, maxSigma: 10, sigmaSigma: 10,
	sigmaAPR: 7, maxLoanAmount: 1000000, minLoanAmount: 1000,
	maxAPR: 19, minAPR: 2.5,
}

type SETTING_TYPE int

const (
	SETTING_INT = iota
	SETTING_FLOAT
	SETTING_STRING
	SETTING_MAX
)

var (
	defaultStocks []stockData = []stockData{
		{Name: "US STEEL"}, {Name: "PAN AM"}, {Name: "FORD"}, {Name: "SANYO"}, {Name: "XEROX"}, {Name: "AT&T"},
	}
	games []gameData

	sr beep.SampleRate
)

const (
	TREND_NONE = iota
	TREND_UP
	TREND_DOWN
	TREND_MAX
)

var trendSymbol [TREND_MAX]string = [TREND_MAX]string{
	"→",
	"↑",
	"↓",
}

func init() {
	sr = beep.SampleRate(44000)
	speaker.Init(sr, 4800)
}
