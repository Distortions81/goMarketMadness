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
	maxPlayers           = 25
	minWeeks             = 2
	maxWeeks             = 52
	maxNameLen           = 64
	startingMoney        = 5000
	randLogarithm        = 100
	minVolatility        = 3.0
	maxVolatility        = 10.0
	volatilityVolatility = 10
	volatilityAPR        = 7
	maxLoanCount         = 10
	maxLoanAmount        = 1000000.0
	minLoanAmount        = 1000
	maxShares            = 10000
	maxAPR               = 19
	minAPR               = 2.5
)

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
