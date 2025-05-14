/*
 * Copyright (C) 2024 Carl Frank Otto III
 * All rights reserved.
 */

package main

import "flag"

func main() {

	skip := flag.Bool("skip", false, "skip startup animations")
	flag.Parse()

	newGame := &gameData{Settings: defSettings}
	go newGame.playGame(*skip)

	startEbiten(newGame)
}
