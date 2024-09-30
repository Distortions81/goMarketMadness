/***************
* STOCK MARKET *
****************/
/*
 * Copyright (C) 2024 Carl Frank Otto III
 * All rights reserved.
 */

package main

import "fmt"

func main() {
	defer fixTerm()
	handleExit()
	setupTerm()

	fmt.Print("\033[2J") //Clear screen
	fmt.Println("Market Madness!")
	fmt.Println("Press any key to begin.")
	anyKey()

	newGame := &gameData{Settings: defSettings}
	newGame.playGame()
}
