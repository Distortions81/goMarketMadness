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
	"log"
	"os"
	"os/exec"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

func init() {
	clearForPlatform = make(map[string]func()) //Initialize it
	clearForPlatform["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clearForPlatform["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	sr = beep.SampleRate(44000)
	err := speaker.Init(sr, 4800)
	if err != nil {
		log.Fatal("speaker init failed.")
	}
}
