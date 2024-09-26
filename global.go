/***************
* STOCK MARKET *
****************/
/*
 * Copyright (C) 2024 Carl Frank Otto III
 * All rights reserved.
 */

package main

import "github.com/faiface/beep"

var (
	clearForPlatform map[string]func()
	stockList        []stockData = []stockData{
		{Name: "US STEEL"}, {Name: "PAN AM"}, {Name: "FORD"}, {Name: "SANYO"}, {Name: "XEROX"}, {Name: "AT&T"},
	}
	numPlayers, numWeeks int
	players              []playerData

	sr beep.SampleRate
)
