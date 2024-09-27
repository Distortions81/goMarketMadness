/***************
* STOCK MARKET *
****************/
/*
 * Inspired by the BASIC game written by
 * Brian Lee and HCM Staff
 * Home computer magazine
 *
 * Copyright (C) 2024 Carl Frank Otto III
 * All rights reserved.
 */

package main

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

func init() {
	sr = beep.SampleRate(44000)
	speaker.Init(sr, 4800)
}
