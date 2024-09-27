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

const (
	maxPlayers       = 10
	minWeeks         = 2
	maxWeeks         = 25
	maxPlayerNameLen = 64
	startingMoney    = 5000

	randLogarithm        = 100
	minVolatility        = 3.0
	maxVolatility        = 10.0
	volatilityVolatility = 10
	volatilityAPR        = 7
	maxLoanCount         = 10
	maxAPR               = 19
	minAPR               = 2.5

	maxLoanSize = 1000000.0
	minLoanSize = 1000
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
