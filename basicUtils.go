/***************
* STOCK MARKET *
****************/
/*
 * Copyright (C) 2024 Carl Frank Otto III
 * All rights reserved.
 */

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/generators"
	"github.com/faiface/beep/speaker"
)

func CallClear() {
	fmt.Print("\033[2J") //Clear screen
}

// Beeps
func CallSound(duration, frequency, volume int) {

	sine, err := generators.SinTone(sr, frequency)
	if err != nil {
		log.Fatal("generators.SinTone: " + err.Error())
	}
	tone := beep.Take(sr.N(time.Duration(duration)*time.Millisecond), sine)
	speaker.Play(tone)
}
