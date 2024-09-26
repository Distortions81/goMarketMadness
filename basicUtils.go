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
	"math/rand"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/generators"
	"github.com/faiface/beep/speaker"
)

func CallClear() {
	fmt.Print("\033[2J") //Clear screen
}

func CallSound(duration, frequency, volume int) {

	sine, err := generators.SinTone(sr, frequency)
	if err != nil {
		log.Fatal("generators.SinTone: " + err.Error())
	}
	tone := beep.Take(sr.N(time.Duration(duration)*time.Millisecond), sine)
	speaker.Play(tone)
}

func RND() float64 {
	return rand.Float64()
}

func DisplayAt(row, col int, data string) {
	row = clamp(row-1, 1, 24)
	col = clamp(col-1, 1, 28)
	fmt.Printf("\033[%d;%dH%v", row, col, data)
}

func SGN(val int) int {
	if val > 0 {
		return 1
	} else if val == 0 {
		return 0
	} else {
		return -1
	}
}
