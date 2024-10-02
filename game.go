/*
 * Copyright (C) 2024 Carl Frank Otto III
 * All rights reserved.
 */

package main

import (
	"time"
)

var gReady bool

func (game *gameData) playGame() {

	countDown := 3
	for x := 0; x < countDown; x++ {
		CallClear()
		printLn("Byte-99/4U")
		//		println("Copyright 2024 Carl Otto III")
		//		println("All rights reserved.")
		printfLn("\nLoading %v...", countDown-x)
		time.Sleep(time.Second)
	}
	CallClear()
	time.Sleep(time.Millisecond * 500)
	printLn("Market Madness!")
	for x := 1; x < 8; x++ {
		CallBGColor(x)
		time.Sleep(time.Millisecond * 200)
	}
	game.setup()

	//Game loop
	for week := range game.NumWeeks {
		game.Week = week + 1

		for p, player := range game.Players {
			if player.Gone {
				continue
			}
			game.showStockPrices()
			printfLn("\nPlayer #%v: %v\nIt is your turn!", p+1, player.Name)
			game.Players[p].processLoans()
			printfLn("Bank balance: $%0.2f", player.Balance)
			promptForChoice(game, player, mainChoiceMenu)
		}

		if game.Week == game.NumWeeks {
			printLn("\n** LAST WEEK!!! ***")
		} else {
			printfLn("\n*** WEEK %v of %v ***", game.Week, game.NumWeeks)
		}
		game.tickStocks()
		game.tickAPR()
	}

	game.showGameStats()

	game.playGame()
}
