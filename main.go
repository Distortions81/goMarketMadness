/*
 * Copyright (C) 2024 Carl Frank Otto III
 * All rights reserved.
 */

package main

func main() {

	newGame := &gameData{Settings: defSettings}
	go newGame.playGame()

	startEbiten(newGame)
}
